<!-- TOC depthTo:1 -->

- [Reference](#reference)
- [6 Rules of Thumb for MongoDB Schema Design](#6-rules-of-thumb-for-mongodb-schema-design)
- [Mongoose with Denormalized Schema](#mongoose-with-denormalized-schema)

<!-- /TOC -->

# Reference

* 6 Rules of Thumb for MongoDB Schema Design
  * [Part 1](https://www.mongodb.com/blog/post/6-rules-of-thumb-for-mongodb-schema-design-part-1)
  * [Part 2](https://www.mongodb.com/blog/post/6-rules-of-thumb-for-mongodb-schema-design-part-2)
  * [Part 3](https://www.mongodb.com/blog/post/6-rules-of-thumb-for-mongodb-schema-design-part-3)
  * [Part 1 zhtw](https://blog.toright.com/posts/4483/mongodb-schema-設計指南.html)
  * [Part 2 zhtw](https://blog.toright.com/posts/4537/mongodb-schema-設計指南-part-ii-反正規化的威力.html)
* Thinking in Documents
  * [Part 1](https://www.mongodb.com/blog/post/thinking-documents-part-1)
  * [Part 2](https://www.mongodb.com/blog/post/thinking-documents-part-2)
  * [Webinar: Back to Basics - Thinking in Documents](https://www.mongodb.com/presentations/webinar-back-to-basics-thinking-in-documents)
    * [slideshare](http://www.slideshare.net/mongodb/webinar-back-to-basics-thinking-in-documents?ref=https://www.mongodb.com/presentations/webinar-back-to-basics-thinking-in-documents)
    <!--* [youtube](https://youtu.be/JohXr_6qWJY)-->

# 6 Rules of Thumb for MongoDB Schema Design

1. favor embedding unless there is a compelling reason not to
1. needing to access an object on its own is a compelling reason not to embed it
1. Arrays should not grow without bound. If there are more than a couple of hundred documents on the “many” side, don’t embed them; if there are more than a few thousand documents on the “many” side, don’t use an array of ObjectID references. High-cardinality arrays are a compelling reason not to embed.
1. Don’t be afraid of application-level joins: if you index correctly and use the projection specifier (as shown in part 2) then application-level joins are barely more expensive than server-side joins in a relational database.
1. Consider the write/read ratio when denormalizing. A field that will mostly be read and only seldom updated is a good candidate for denormalization: if you denormalize a field that is updated frequently then the extra work of finding and updating all the instances is likely to overwhelm the savings that you get from denormalizing.
1. As always with MongoDB, how you model your data depends – entirely – on your particular application’s data access patterns. You want to structure your data to match the ways that your application queries and updates it.


<!-- TOC depthFrom:2 -->

- [One-to-Few](#one-to-few)
- [One-to-Many](#one-to-many)
- [One-to-Squillions](#one-to-squillions)
- [Two-Way Referencing](#two-way-referencing)
- [Denormalizing](#denormalizing)
  - [One-To-Many Relationships](#one-to-many-relationships)
    - [Denormalizing from Many -> One](#denormalizing-from-many---one)
    - [Denormalizing from One -> Many](#denormalizing-from-one---many)
  - [One-To-Squillions Relationships](#one-to-squillions-relationships)

<!-- /TOC -->

## One-to-Few

* Embed the N side if the cardinality is one-to-few and there is no need to access the embedded object outside the context of the parent object

This design has all of the advantages and disadvantages of embedding.

* **Advantage** is that you don't have to perform a separate query to get the embedded details;
* **Disadvantage** is that you have no way of accessing the embedded details as stand-alone entities.

```js
> db.person.findOne()
{
  name: 'Kate Monster',
  ssn: '123-456-7890',
  addresses : [
     {street: '123 Sesame St', city: 'Anytown', cc: 'USA'},
     {street: '123 Avenue Q', city: 'New York', cc: 'USA'}
  ]
}
```

## One-to-Many

* Use an array of references to the N-side objects if the cardinality is one-to-many or if the N-side objects should stand alone for any reasons

This style of referencing has a complementary set of advantages and disadvantages to embedding.

* **Advantage** Each Part is a stand-alone document, so it's easy to search them and update them independently.
* **Disadvantage** One trade off for using this schema is having to perform a second query to get details about the Parts for a Product.

```js
> db.parts.findOne()
{
    _id : ObjectID('AAAA'),
    partno : '123-aff-456',
    name : '#4 grommet',
    qty: 94,
    cost: 0.94,
    price: 3.99
}

> db.products.findOne()
{
    name : 'left-handed smoke shifter',
    manufacturer : 'Acme Corp',
    catalog_number: 1234,
    parts : [     // array of references to Part documents
        ObjectID('AAAA'),    // reference to the #4 grommet above
        ObjectID('F17C'),    // reference to a different Part
        ObjectID('D2AA'),
        // etc
    ]
}
```

You would then use an application-level join to retrieve the parts for a particular product:

```js
// Fetch the Product document identified by this catalog number
> product = db.products.findOne({catalog_number: 1234});
// Fetch all the Parts that are linked to this Product
> product_parts = db.parts.find({_id: { $in : product.parts} } ).toArray() ;
```

## One-to-Squillions

* Use a reference to the One-side in the N-side objects if the cardinality is one-to-squillions

Any given host could generate enough messages to overflow the `16 MB` document size, even if all you stored in the array was the ObjectID.

This is the classic use case for `parent-referencing` – you’d have a document for the host, and then store the ObjectID of the host in the documents for the log messages.

```js
> db.hosts.findOne()
{
    _id : ObjectID('AAAB'),
    name : 'goofy.example.com',
    ipaddr : '127.66.66.66'
}

> db.logmsg.findOne()
{
    time : ISODate("2014-03-28T09:42:41.382Z"),
    message : 'cpu is on fire!',
    host: ObjectID('AAAB')       // Reference to the Host document
}
```

## Two-Way Referencing

This design has all of the advantages and disadvantages of the “One-to-Many” schema, but with some additions.

using this schema design means that it is no longer possible to reassign a Task to a new Person with a single atomic update.

```js
db.person.findOne()
{
    _id: ObjectID("AAF1"),
    name: "Kate Monster",
    tasks [     // array of references to Task documents
        ObjectID("ADF9"),
        ObjectID("AE02"),
        ObjectID("AE73")
        // etc
    ]
}

db.tasks.findOne()
{
    _id: ObjectID("ADF9"),
    description: "Write lesson plan",
    due_date:  ISODate("2014-04-01"),
    owner: ObjectID("AAF1")     // Reference to Person document
}
```

## Denormalizing

* Denormalizing saves you a lookup of the denormalized data at the cost of a more expensive update.
* Denormalizing only makes sense when there’s an high ratio of reads to updates.
* Also note that if you denormalize a field, you lose the ability to perform atomic and isolated updates on that field. Just like with the two-way referencing example above.

### One-To-Many Relationships

#### Denormalizing from Many -> One

```js
> db.products.findOne()
{
    name : 'left-handed smoke shifter',
    manufacturer : 'Acme Corp',
    catalog_number: 1234,
    parts : [
        {id : ObjectID('AAAA'), name : '#4 grommet' },         // Part name is denormalized
        {id: ObjectID('F17C'), name : 'fan blade assembly' },
        {id: ObjectID('D2AA'), name : 'power switch' },
        // etc
    ]
}
```

While making it easier to get the part names, this would add just a bit of client-side work to the application-level join:

```js
// Fetch the product document
> product = db.products.findOne({catalog_number: 1234});
  // Create an array of ObjectID()s containing *just* the part numbers
> part_ids = product.parts.map(function(doc) { return doc.id } );
  // Fetch all the Parts that are linked to this Product
> product_parts = db.parts.find({_id: { $in : part_ids} } ).toArray() ;
```

#### Denormalizing from One -> Many

```js
> db.parts.findOne()
{
    _id : ObjectID('AAAA'),
    partno : '123-aff-456',
    name : '#4 grommet',
    product_name : 'left-handed smoke shifter',   // Denormalized from the ‘Product’ document
    product_catalog_number: 1234,                     // Ditto
    qty: 94,
    cost: 0.94,
    price: 3.99
}
```

### One-To-Squillions Relationships

This works in one of two ways:
* put information about the “one” side (from the 'hosts’ document) into the “squillions” side (the log entries).
* put summary information from the “squillions” side into the “one” side.

```js
> db.logmsg.findOne()
{
    time : ISODate("2014-03-28T09:42:41.382Z"),
    message : 'cpu is on fire!',
    ipaddr : '127.66.66.66',  // from hosts
    host: ObjectID('AAAB')
}

// denormalize it ALL into the “squillions” side and get rid of the “one” collection altogether:
> db.logmsg.findOne()
{
    time : ISODate("2014-03-28T09:42:41.382Z"),
    message : 'cpu is on fire!',
    ipaddr : '127.66.66.66',        // from hosts
    hostname : 'goofy.example.com', // from hosts
}
```

On the other hand, you can also denormalize into the “one” side.  Lets say you want to keep the last 1000 messages from a host in the 'hosts’ document.

```js
//  Get log message from monitoring system
logmsg = get_log_msg();
log_message_here = logmsg.msg;
log_ip = logmsg.ipaddr;
// Get current timestamp
now = new Date()
// Find the _id for the host I’m updating
host_doc = db.hosts.findOne({ipaddr : log_ip},{_id:1});  // Don’t return the whole document
host_id = host_doc._id;
// Insert the log message, the parent reference, and the denormalized data into the ‘many’ side
db.logmsg.save({time : now, message : log_message_here, ipaddr : log_ip, host : host_id) });
// Push the denormalized log message onto the ‘one’ side
db.hosts.update({_id: host_id},
  {$push : {logmsgs : { $each:  [ { time : now, message : log_message_here} ],
                      $sort:  {time : 1},  // Only keep the latest ones
                      $slice: -1000 }        // Only keep the latest 1000
  }} );
```


# Mongoose with Denormalized Schema

reference: [mongoose issue 4736](https://github.com/Automattic/mongoose/issues/4736).

```js
// schema
var personSchema = new Schema({
  name: String,
  email: String
});

var subscriptionSchema = new Schema({
  subPerson: personSubSchema
});

// sub scheam
var personSubSchema = new Schema({
  name: String
});

// virtual schema
var virtual = subscriptionSchema.virtual('person', {
  ref: 'Person',
  localField: 'subPerson._id',
  foreignField: '_id',
  justOne: true
});

// getters are executed in reverse order, so this executes after the default getter for virtual populate
// parameter to getter is the value returned from the previous getter, so if the virtual has a value then v will be defined, otherwise falsy
virtual.getters.unshift(function(v) {
  return v || this.subPerson;
});

var Person = mongoose.model('Person', personSchema);
var Subscription = mongoose.model('Subscription', subscriptionSchema);

Person.create({ name: 'Val', email: 'a@b.co' }).
  then(val => Subscription.create({ person: val })).
  then(sub => Promise.all([
    Subscription.findById(sub).populate('person'),
    Subscription.findById(sub)
  ])).
  then(res => {
    res.forEach(res => {
      console.log(res.toObject({ virtuals: true }));
    })
  }).
  catch(error => {
    console.error(error);
    process.exit(-1);
  });
```

log result:

```js
// unpopulated document
{
  _id: ...,
  person: { _id: ..., name: 'Val' },
  subPerson: { _id: ..., name: 'Val' }
}
// populated document
{
  _id: ...,
  person: { _id: ...,  name: 'Val', email: 'a@b.co' }
  subPerson: { _id: ..., name: 'Val' }
}
```
