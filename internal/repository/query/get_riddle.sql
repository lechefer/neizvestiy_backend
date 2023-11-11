select r.id,
       r.quest_step_id,
       r.name,
       r.description,
       r.status,
       r.letter,
       coalesce(ar.riddle_status, 'not_passed') as status
from riddles as r
         left join account_riddles as ar on ar.riddle_id = r.id and ar.account_id = $1
where r.id = $2;