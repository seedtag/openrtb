## Seedtag's OpenRTB

This library is forked as a close copy of previd/openrtb but with neccesary changes to operate at Seedtag.




## What's different from prebid/openrtb

### default=0 and omitempty

At prebid/openrtb, any field that is integer with a default=0, is created with the `omitempty` property. The main reason to do that is saving some bytes on the transferred payload.

However, at seedtag we think that some partners could have implemented openrtb in a way that this undefined values could be taken differently as 0 values, mostly for key fields as `imp.bidfloor`.

The list of fields where we have removed the `omitempty` property:
- req.imp.bidfloor
- req.imp.instl
- req.test
- req.allimps

### coppa

Despiste not having a default, the `coppa` field is handled as `int` and `omitempty` on the official prebid/openrtb.

We've moved this type to `*int` to be able to:
- use nil when we don't want to add any info (omitempty will do teh rest)
- use pointer to 0 when we want to explicitely say NO
- use pointer to 1 when we want to explicitely say YES

