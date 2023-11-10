select a.id, a.name, a.icon, a.steps, a.description, coalesce(aa.passed,0) as passed, coalesce(aa.is_completed, false) as is_completed
from achievements as a
         left join accounts_achievements as aa on aa.achievement_id = a.id and aa.account_id = $1;
