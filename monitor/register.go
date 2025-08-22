package monitor

import (
	"github.com/anonymous-author-00000/ra-webs/monitor/ent"
	"github.com/anonymous-author-00000/ra-webs/monitor/serviceclient"
)

func (monitor *Monitor) RegisterTA(publicKey []byte) *ent.TA {
	ta := monitor.DB.Client.TA.
		Create().
		SetPublicKey(publicKey).
		SaveX(*monitor.DB.Ctx)

	return ta
}

func (monitor *Monitor) RegisterCTLog(ctLogId int, ta *ent.TA) *ent.CTLog {
	ctLog := monitor.DB.Client.CTLog.
		Create().
		SetMonitorLogID(ctLogId).
		SetTa(ta).
		SaveX(*monitor.DB.Ctx)

	return ctLog
}

func (monitor *Monitor) RegisterEvidenceLog(uniqueId []byte, entry *serviceclient.EvidenceEntry, ta *ent.TA) *ent.EvidenceLog {
	evidenceLog := monitor.DB.Client.EvidenceLog.
		Create().
		SetEvidence(entry.Evidence).
		SetRepository(entry.Repository).
		SetCommitID(entry.CommitID).
		SetUniqueID(uniqueId).
		SetTa(ta).
		SaveX(*monitor.DB.Ctx)

	return evidenceLog
}
