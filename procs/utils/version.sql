CREATE EXTENSION plpython3u;
CREATE FUNCTION get_alibi_detect_version() RETURNS text AS $$
import alibi_detect
return alibi_detect.__version__
$$ LANGUAGE plpython3u;