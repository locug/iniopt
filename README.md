# iniopt
Compare two INI files creating a SQL file for updating INI files from samples


## INFO
### What created file sould contain, windows 1252 with CRLF line endings
Set one or many settings - INIFILENAME must contain ".INI"
@WIZSET(<<INIFILENAME>>[<<SECTIONNAME>>]<<KEY>>=<<VALUE>>);

Write to the INI file
@dbHot(<<INIFILENAME>>,SET,<<INIFILENAME>>[*);
