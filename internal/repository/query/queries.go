package query

import _ "embed"

// Quest

//go:embed get_quest.sql
var GetQuestSql string

//go:embed get_quest_steps.sql
var GetQuestStepsSql string

//go:embed list_quests.sql
var ListQuestsSql string

// Settlements

//go:embed list_settlements.sql
var ListSettlementsSql string

//go:embed get_achievement.sql
var GetAchievementSql string

//go:embed list_achievements.sql
var ListAchievementsSql string
