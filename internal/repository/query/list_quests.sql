select q.*
from quests as q
where q.settlement_id = $1;