CREATE FUNCTION isSenior(e individuals) RETURNS boolean AS $$
  if e["age"] > 55:
    return True
  if (e["age"] < 55) and (e["age"] > 45):
    return True
  return False
$$ LANGUAGE plpython3u;

SELECT *, isSenior(individuals) as senior FROM individuals WHERE isSenior(individuals) is TRUE;