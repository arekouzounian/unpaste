package cmd

/*
	NOTE: might want to look into JSON lines format
	or some other format that allows for storing
	data in a formatted, easily read way, but also
	is easy to append to. As more flags are added
	(such as the aggregate-data flag), I'm going to want
	to store more data into these types, so it's easiest
	to use JSON to serialize and deserialize the data back and
	forth, but keep in mind that as we get to larger and larger
	datasets, deserializing the entire file into memory will
	be terribly inefficient and not safe in the event of a
	file leak.
*/

type ArchiveLink struct {
	Key            string `json:"key"`
	GroupDirectory string `json:"group_dir"`
}

type ArchiveLinkEntry struct {
	Number int           `json:"index"`
	List   []ArchiveLink `json:"list"`
}
