update account_quest_steps
set status = 'ended'
where account_id = $1 and quest_step_id = $2;