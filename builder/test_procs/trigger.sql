CREATE FUNCTION update_row() RETURNS trigger AS $$
import alibi_detect
TD["new"]["age"] = 10
$$ LANGUAGE plpython3u;


CREATE TRIGGER udtRow BEFORE INSERT ON individuals
    FOR EACH ROW EXECUTE FUNCTION update_row();