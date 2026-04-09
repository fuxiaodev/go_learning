package L3
/*
slices wrap arrays to give more general, powerful, and convenient interface to
sequences of data. Except for items with explicit dimensions such as transformation matrices.
most programming in go is done with slices rather than simple arrays

Slices hold references to an underlying array, if you assign one slice to another, bother
refer to the same array.
if a function takes a slice argument, any changes it makes to the elements of the slice will be visible
to the caller, analogous to passing a pointer to the underlying array
*/