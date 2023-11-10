with result as (select s.id,
                       s.name,
                       s.location,
                       case
                           when $1 = '' then 0.0
                           else word_similarity(s.name, $1)
                           end as sml
                from settlements as s
                where $1 = ''
                   OR $1 <% s.name
                order by sml desc
                offset $2 limit $3)
select r.id, r.name, r.location
from result as r
order by r.sml desc;