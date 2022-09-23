CREATE FUNCTION exec_raw() RETURNS BOOLEAN AS $$

detect_function = "CREATE FUNCTION test_fn(e text) RETURNS TEXT AS && return 'abc' && LANGUAGE plpython3u;"
detect_function = detect_function.replace('&','$')
plpy.info(detect_function)
return True

$$ LANGUAGE plpython3u;

SELECT * FROM exec_raw();

DROP FUNCTION exec_raw;