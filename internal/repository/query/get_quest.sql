select q.id,
       q.settlement_id,
       q.name,
       q.description,
       q.type,
       q.avg_duration,
       q.reward,
       coalesce(aq.is_active, false) as is_active
from quests as q
         left join account_quests as aq on aq.quest_id = q.id and aq.account_id = $1
where q.id = $2;