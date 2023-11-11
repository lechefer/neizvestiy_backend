package query

import _ "embed"

// Account

//go:embed create_account.sql
var CreateAccountSql string

// Quest

//go:embed get_quest.sql
var GetQuestSql string

//go:embed get_quest_steps.sql
var GetQuestStepsSql string

//go:embed start_quest.sql
var StartQuestSql string

//go:embed end_quest_step.sql
var EndQuestStepSql string

// Settlements

//go:embed list_quests.sql
var ListQuestsSql string

//go:embed list_settlements.sql
var ListSettlementsSql string

// Achievements

//go:embed get_achievement.sql
var GetAchievementSql string

//go:embed list_achievements.sql
var ListAchievementsSql string

// Riddles

//go:embed list_riddles.sql
var ListRiddlesSql string

//go:embed update_riddle.sql
var UpdateRiddleSql string

//go:embed get_riddle.sql
var GetRiddleSql string
