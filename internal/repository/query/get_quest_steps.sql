select qs.id,
       qs.quest_id,
       qs."order",
       qs.name,
       qs.place_type,
       qs.address,
       qs.phone,
       qs.email,
       qs.website,
       qs.schedule,
       qs.location,
       qs.status,
       coalesce(aqs.status, 'inactive') as status
from quests_steps as qs
left join account_quest_steps as aqs on aqs.quest_step_id = qs.id and aqs.account_id = $1
where qs.quest_id = $2;