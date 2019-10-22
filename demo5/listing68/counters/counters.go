// counters 包提供告警计数器的功能
package counters

//alertCounter 未公开的锁
//保存警告计数
type alertCounter int

//New返回一个未公开的 alertCounter值
func New(value int) alertCounter {
	return alertCounter(value)
}
