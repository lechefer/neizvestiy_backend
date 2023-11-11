select r.*
from riddles as r
where r.quest_step_id = $1;