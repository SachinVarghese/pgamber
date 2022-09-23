CREATE FUNCTION trainVAEDetector(table_name text, outlier_perc int) RETURNS TEXT AS $$

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

filepath = '/var/lib/postgresql/data/detectors/outlier/'+table_name

save_detector(od, filepath)
plpy.info("Saved outlier detector at:")
return filepath;

$$ LANGUAGE plpython3u;

SELECT * FROM trainVAEDetector('individuals',10);

DROP FUNCTION trainVAEDetector;