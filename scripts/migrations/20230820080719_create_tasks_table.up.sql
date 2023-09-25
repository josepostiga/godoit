create table if not exists tasks (
    id int primary key generated always as identity,
    title text not null,
    description text,
    created_at timestamptz null default current_timestamp,
    updated_at timestamptz null default current_timestamp,
    completed_at timestamptz null
);
