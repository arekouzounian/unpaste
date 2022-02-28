package cmd

type ArchiveLink struct {
	Key            string `json:"key"`
	GroupDirectory string `json:"group_dir"`
}

type ArchiveLinkEntry struct {
	Number int           `json:"index"`
	List   []ArchiveLink `json:"list"`
}
