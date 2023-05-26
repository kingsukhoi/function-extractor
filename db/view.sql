CREATE EXTENSION pg_trgm;


create view similarity(func_name, base_path, pay_path, similarity) as
SELECT base.func_name,
       base.relpath                                                 AS base_path,
       paybotic.relpath                                             AS pay_path,
       similarity(base.body, paybotic.body) * 100::double precision AS similarity
FROM (SELECT functions.abspath,
             functions.func_name,
             functions.body,
             functions.type,
             functions.id,
             functions.relpath
      FROM functions
      WHERE functions.type = 'base'::text) base
         JOIN (SELECT functions.abspath,
                      functions.func_name,
                      functions.body,
                      functions.type,
                      functions.id,
                      functions.relpath
               FROM functions
               WHERE functions.type = 'paybotic'::text) paybotic
              ON paybotic.relpath = base.relpath AND paybotic.func_name = base.func_name;


