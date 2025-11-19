//go:generate ../bin/go-enum --marshal --lower -b example

package example

// RiskLevel is an enumeration of risk levels.
/*
ENUM(
normal `ZH:"无风险" RU:"Риска риск"`,   // No risk
low    `ZH:"低风险" RU:"низкий риск"`   // Low risk
medium `ZH:"中风险" RU:"Средний риск"`, // Medium risk
high   `ZH:"高风险" RU:"Высокий риск"`  // High risk
)
*/
type RiskLevel int8
