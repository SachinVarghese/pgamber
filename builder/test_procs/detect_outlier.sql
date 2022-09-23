CREATE FUNCTION isOutlier(e individuals, excluded_features int[])RETURNS BOOLEAN AS $$

import numpy as np
from alibi_detect.utils.saving import load_detector

if not "od" in SD:
  filepath = "/var/lib/postgresql/data/detectors/outlier/"+"individuals"
  od = load_detector(filepath)
  SD["od"] = od
else:
  od = SD["od"]

if not "mu" in SD:
  plan = plpy.prepare("SELECT * FROM individuals;", [])
  rv = plpy.execute(plan, [])
  features = rv.colnames()
  rowNum = rv.nrows()

  X_ref = np.zeros((rowNum, len(features)))
  for i in range(rowNum):
    for j in range(len(features)):
      X_ref[i][j] = rv[i][features[j]]


  X_ref = np.delete(X_ref,excluded_features,1) 

  mu, sigma = X_ref.mean(axis=0), X_ref.std(axis=0)
  SD["mu"] = mu
  SD["sigma"] = sigma
else:
  mu = SD["mu"]
  sigma = SD["sigma"]

features = list(e.keys())
X_test = np.zeros((1, len(features)))
for i in range(len(features)):
  X_test[0][i] = e[features[i]]

X_test = np.delete(X_test,excluded_features,1)  
X_test = (X_test - mu) / sigma

preds = od.predict(X_test,outlier_type='instance', return_feature_score=False, return_instance_score=False)

return preds['data']['is_outlier']>0

$$ LANGUAGE plpython3u;

SELECT *, isOutlier(individuals,ARRAY[0]) as outlier FROM individuals LIMIT 20;

SELECT *, isOutlier(individuals) as outlier FROM individuals WHERE age > 55;

SELECT *, isOutlier(individuals) as outlier FROM individuals WHERE isOutlier(individuals) is TRUE;

DROP FUNCTION isOutlier;