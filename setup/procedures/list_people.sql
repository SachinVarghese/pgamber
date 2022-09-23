CREATE FUNCTION list_individuals() RETURNS SETOF individuals AS $$
import alibi_detect
plan = plpy.prepare("SELECT * FROM individuals WHERE age > $1;", ["int"])
rv = plpy.execute(plan, [50], 5)
return rv
$$ LANGUAGE plpython3u;

SELECT * FROM list_individuals();