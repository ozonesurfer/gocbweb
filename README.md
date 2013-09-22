# Database
You will need to install Couchbase. You can find the download instructions at [www.couchbase.com](www.couchbase.com). After installing it, you will need to set up a Bucket named "gocbweb". You will also need to create a development view called "dev_gocbweb" with the following views:

"all_band":

> function (doc, meta) {

>> if (doc.Type && doc.Type == "band") {

>>> emit(null, doc);

>> }

> } 

"all_genre":

> function (doc, meta) {

>> if (doc.Type && doc.Type == "genre") {

>>> emit(null, doc);

>> }

> }

"all_location":

> function (doc, meta) {

>> if (doc.Type && doc.Type == "location") {

>>> emit(null, doc);

>> }

> }

"by_genre":

> function (doc, meta) {

>> if (doc.Value.Albums) {

>>> for (i = 0; i < doc.Value.Albums.length; i++) {

>>>> emit(doc.Value.Albums[i].GenreId, doc);
 
>>> }
 
>> }
 
> }

# Dependencies

You will need to run the "go install" command to install the following Go packages:

+ github.com/couchbaselabs/go-couchbase
+ github.com/QLeelulu/goku

The Goku package is used to configure the framework to construct MVC websites.

# Configuration

You will need to change the Couchbase bucket password, Couchbase pool URL, and website URL/port number in /src/gocbweb/config.go to match your server's configuration. 