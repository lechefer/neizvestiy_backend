select q.*,
    coalesce(aq.is_active, false) as is_active
from quests as q
    left join account_quests as aq on aq.quest_id = q.id and aq.account_id = $1
where q.settlement_id = $2;