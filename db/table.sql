create table functions
(
    abspath   text not null,
    func_name text not null,
    body      text not null,
    type      text,
    id        uuid default uuid_generate_v4(),
    relpath   text
);