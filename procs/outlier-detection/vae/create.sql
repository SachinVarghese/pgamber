CREATE EXTENSION plpython3u;
CREATE FUNCTION trainVAEOutlierDetector(table_name text, excluded_features int[], outlier_perc int) RETURNS TEXT AS $$

import numpy as np
import tensorflow as tf
tf.keras.backend.clear_session()
from tensorflow.keras.layers import Dense, InputLayer
from alibi_detect.od import OutlierVAE
from alibi_detect.utils.saving import save_detector

plpy.info("Creating outlier detector for dataset: ",table_name)
plan = plpy.prepare("SELECT * FROM "+table_name+";", [])
rv = plpy.execute(plan, [])

features = rv.colnames()
rowNum = rv.nrows()

X_ref = np.zeros((rowNum, len(features)))
for i in range(rowNum):
  for j in range(len(features)):
    X_ref[i][j] = rv[i][features[j]]

X_ref = np.delete(X_ref,excluded_features,1)    

mu, sigma = X_ref.mean(axis=0), X_ref.std(axis=0)
X_ref = (X_ref - mu) / sigma

n_features = X_ref.shape[1]
latent_dim = 2

encoder_net = tf.keras.Sequential(
  [
      InputLayer(input_shape=(n_features,)),
      Dense(25, activation=tf.nn.relu),
      Dense(10, activation=tf.nn.relu),
      Dense(5, activation=tf.nn.relu)
  ])

decoder_net = tf.keras.Sequential(
  [
      InputLayer(input_shape=(latent_dim,)),
      Dense(5, activation=tf.nn.relu),
      Dense(10, activation=tf.nn.relu),
      Dense(25, activation=tf.nn.relu),
      Dense(n_features, activation=None)
  ])

od = OutlierVAE(threshold=None, score_type='mse', encoder_net=encoder_net,  decoder_net=decoder_net, latent_dim=latent_dim, samples=5)
od.fit(X_ref,loss_fn=tf.keras.losses.mse,epochs=5,verbose=False)
od.infer_threshold(X_ref, threshold_perc=100-outlier_perc, outlier_perc=100)

filepath = '/var/lib/postgresql/data/detectors/outlier/vae/'+table_name
save_detector(od, filepath)

detect_function = """
CREATE FUNCTION isVAEOutlier(record $1) 
RETURNS BOOLEAN AS && 
import numpy as np
from alibi_detect.utils.saving import load_detector

if not 'od' in SD:
  filepath = '$2'
  od = load_detector(filepath)
  SD['od'] = od
else:
  od = SD['od']

mu = $3
sigma = $4

features = list(record.keys())
X_test = np.zeros((1, len(features)))
for i in range(len(features)):
  X_test[0][i] = record[features[i]]

X_test = np.delete(X_test,$5,1)  
X_test = (X_test - mu) / sigma

preds = od.predict(X_test,outlier_type='instance', return_feature_score=False, return_instance_score=False)

return preds['data']['is_outlier']>0
&& LANGUAGE plpython3u;
"""

detect_function = detect_function.replace('$1',table_name)
detect_function = detect_function.replace('$2',filepath)
detect_function = detect_function.replace('$3',np.array2string(mu,separator=','))
detect_function = detect_function.replace('$4',np.array2string(sigma,separator=','))
detect_function = detect_function.replace('$5',np.array2string(np.array(excluded_features),separator=','))
detect_function = detect_function.replace('&','$')

plan = plpy.prepare(detect_function, [])
rv = plpy.execute(plan, [])

plpy.info("Saved outlier detector and created detection function: isVAEOutlier")
return filepath;

$$ LANGUAGE plpython3u;


CREATE FUNCTION dropVAEOutlierDetector(table_name text) RETURNS TEXT AS $$
import shutil
filepath = '/var/lib/postgresql/data/detectors/outlier/vae/'+table_name
shutil.rmtree(filepath)

plan = plpy.prepare("DROP FUNCTION isVAEOutlier("+table_name+");", [])
rv = plpy.execute(plan, [])
return "removed vae outlier detector for table: "+ table_name
$$ LANGUAGE plpython3u;