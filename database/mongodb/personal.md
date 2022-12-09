# Reading notes:
- Experience index
  + Query multiple fields -> using compound index
  + The order of fields in the index really important
    + By exact > range > sort query
    + By diversity many > little data
    + By frequency many > little query
  + Using partital index with big collections
  + Using hashed index with longest field need search exactly (ex: token, uuid)
<h5></h5>

- Optimize application
  + 99% operation backend has bottleneck at DB -> Don't using DB to calculate. Should using backend 
  + Optimize UX/UI -> Every user interaction = one query into 1 collection
  + Design filter & search exactly when inititial database
  + Big collections shouldn't query prefix, regex, text -> Using elasticsearch
  + Caching data which little change.