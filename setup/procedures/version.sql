CREATE FUNCTION get_alibi_version() RETURNS text AS $$
import alibi_detect
return alibi_detect.__version__
$$ LANGUAGE plpython3u;

SELECT get_alibi_version();