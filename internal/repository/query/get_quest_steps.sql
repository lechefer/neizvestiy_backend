select qs.*
from quests_steps as qs
where qs.quest_id = $1;