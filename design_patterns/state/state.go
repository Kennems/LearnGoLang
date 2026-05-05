package state
import "strconv"

type Work struct {
	hour  int
	state State
}
func NewWork() *Work { return &Work{state: &MoringState{}} }
func (w *Work) Hour() int { return w.hour }
func (w *Work) SetHour(h int) { w.hour = h }
func (w *Work) SetState(s State) { w.state = s }
func (w *Work) WriteProgram() string { return w.state.WriteProgram(w) }

type State interface {
	WriteProgram(w *Work) string
}

type MoringState struct{}
func (m *MoringState) WriteProgram(w *Work) string {
	if w.Hour() < 12 { return "上午工作，精神百倍:" + strconv.Itoa(w.Hour()) }
	w.SetState(&NoonState{})
	return w.WriteProgram()
}

type NoonState struct{}
func (n *NoonState) WriteProgram(w *Work) string {
	if w.Hour() < 13 { return "饿了，午饭；犯困，午休:" + strconv.Itoa(w.Hour()) }
	w.SetState(&AfternoonState{})
	return w.WriteProgram()
}

type AfternoonState struct{}
func (a *AfternoonState) WriteProgram(w *Work) string {
	if w.Hour() < 17 { return "下午状态还不错，继续努力:" + strconv.Itoa(w.Hour()) }
	w.SetState(&EveningState{})
	return w.WriteProgram()
}

type EveningState struct{}
func (e *EveningState) WriteProgram(w *Work) string {
	if w.Hour() < 21 { return "加班哦，疲累之极:" + strconv.Itoa(w.Hour()) }
	return "不行了，睡着了:" + strconv.Itoa(w.Hour())
}
