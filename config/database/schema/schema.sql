drop database if exists sentinel;
create database sentinel;
use sentinel;

/*
 - the database has the following functional areas
   1. user
   2. systems
   3. 
*/


-- the user section
create table user (
    id int not null auto_increment,

    display_name varchar(64) not null unique, -- the users unique display name
    email varchar(256) not null unique,
    password varchar(256) not null,
    enabled boolean not null default false, -- indicates if the user account is enabled, defaults to not enabled: false
    roles varchar(256) not null default '["user"]', -- user roles allocated by the application
    last_time_logged_in timestamp not null default now(), -- a timestamp indicating the last time the user logged in
    log_in_count smallint not null default 0, -- a count of each successful login

    updated timestamp not null default now() on update now(),
    created timestamp not null default now(),
    
    primary key(id)
);

/*
 - table functional requirement
 - what high level function must the asset perform, why does the asset exist
*/

create table functional_requirement (
    id int not null auto_increment, -- primary key

    title varchar(64) not null, -- the formal title of the functional requirement
    statement varchar(512) not null, -- holds the functional requirement statement
    reference varchar(64) not null, -- holds an external requirements reference
    reference_source varchar(16) not null, -- legislation, standard, contract, best practice 

    updated timestamp not null default now() on update now(),
    created timestamp not null default now(),

    primary key (id),
    constraint ck_functional_requirement__reference_source
        check (reference_source in ('Legislation', 'Standard', 'Contract', 'Best Practice'))
);

create table functional_system (
    id int not null auto_increment, -- primary key
    functional_requirement_ref int not null, -- fk reference to the functional requirements table

    name varchar(64) not null, -- the name of the functional system
    description varchar(512) not null, -- more detail about the functional system

    updated timestamp not null default now() on update now(),
    created timestamp not null default now(),

    primary key (id),
    constraint fk_functional_system__functional_requirement
        foreign key (functional_requirement_ref) references functional_requirement (id) 
        on update cascade 
        on delete cascade
);

create table asset (
    id int not null auto_increment, -- primary key
    functional_system_ref int not null, -- fk reference to the functional system table

    name varchar(64) not null, -- the name of the asset
    description varchar(512) not null, -- more detail about the asset

    updated timestamp not null default now() on update now(),
    created timestamp not null default now(),

    primary key (id),
    constraint fk_asset__functional_system
        foreign key (functional_system_ref) references functional_system (id)
        on update cascade
        on delete cascade
);

create table asset_part (
    id int not null auto_increment, -- primary key
    asset_ref int not null, -- fk reference to the asset table

    name varchar(64) not null, -- what name has the manufacturer given this part

    updated timestamp not null default now() on update now(),
    created timestamp not null default now(),

    primary key (id),
    constraint fk_asset_part__asset
        foreign key (asset_ref) references asset (id)
        on update cascade
        on delete cascade
);
