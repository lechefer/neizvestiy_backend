-- +goose Up
-- Settlements

insert into "accounts"(id)
values('1'),
      ('2'),
      ('3');

insert into "settlements"("id", "name", "location")
values ('f218690b-4948-4217-8c76-7a433f533f42', 'Смоленск', point(54.785015, 32.043574)),
       ('a5e5ea94-e176-433f-9dac-b9840a039fe0', 'Пенза', point(53.195042, 45.018316));

-- Quests
insert into "quests"("id", "settlement_id", "name", "description", "type", "avg_duration", "reward")
values ('dbb2edca-b8fb-4d9d-bb65-1461a6fed8c9', 'f218690b-4948-4217-8c76-7a433f533f42', 'Смоленск - щит России',
        'Этот маршрут предлагает погрузиться в историю и героическое прошлое города Смоленск. Смоленск является одним из самых важных исторических городов России, богатым событиями и памятниками культуры. Город занимает значительное место в российской и мировой истории благодаря своей стратегической важности, оборонительным сражениям и культурному наследию.\n\nОдним из наиболее известных событий, связанных с Смоленском, является Смоленская битва 1812 года, в рамках войн против Наполеона. В этом сражении русские войска под командованием Михаила Кутузова сражались с французской армией, что сыграло ключевую роль в остановке наступления Наполеона на Россию.',
        'route', extract(epoch from '2 hours 30 minutes'::interval) * 1000000000, '60'),
       
insert into "quests_steps"("id", "quest_id", "order", "name", "place_type", "address", "phone", "email", "website", "schedule", "location", "status")
values ('51219208-8c66-4320-ac02-d37f7998ef63', 'dbb2edca-b8fb-4d9d-bb65-1461a6fed8c9', 1, 'Русская старина', 'Музей', 'ул. Тенишевой 7, Смоленск', '+74812339585', 'museum@smolkrepost.ru', 'https://smolkrepost.ru', '[{"WeekDay":"Tue","FromTo":"10:00 – 18:00"},{"WeekDay":"Wed","FromTo":"10:00 – 18:00"},{"WeekDay":"Thu","FromTo":"10:00 – 18:00"},{"WeekDay":"Fri","FromTo":"10:00 – 17:00"},{"WeekDay":"Sat","FromTo":"10:00 – 18:00"},{"WeekDay":"Sun","FromTo":"10:00 – 18:00"}]', point(54.776799, 32.053039), 'inactive'),
       ('d297d769-cab7-4604-93a5-64101fe0df42', 'dbb2edca-b8fb-4d9d-bb65-1461a6fed8c9', 2, 'Смоленск - щит России', 'Музей', 'ул. Октябрьской Революции 3, Смоленск', '+74812339585', 'museum@smolkrepost.ru', 'https://smolkrepost.ru', '[{"WeekDay":"Tue","FromTo":"10:00 – 18:00"},{"WeekDay":"Wed","FromTo":"10:00 – 18:00"},{"WeekDay":"Thu","FromTo":"10:00 – 18:00"},{"WeekDay":"Fri","FromTo":"10:00 – 17:00"},{"WeekDay":"Sat","FromTo":"10:00 – 18:00"},{"WeekDay":"Sun","FromTo":"10:00 – 18:00"}]', point(54.777941, 32.044505),'active'),
       ('6c35331c-9df1-4257-8b91-1be0e355bc2c', 'dbb2edca-b8fb-4d9d-bb65-1461a6fed8c9', 3, 'Исторический музей', 'Музей', 'ул. Ленина 8, Смоленск', '+74812339585', 'museum@smolkrepost.ru', 'https://smolkrepost.ru', '[{"WeekDay":"Tue","FromTo":"10:00 – 18:00"},{"WeekDay":"Wed","FromTo":"10:00 – 18:00"},{"WeekDay":"Thu","FromTo":"10:00 – 18:00"},{"WeekDay":"Fri","FromTo":"10:00 – 17:00"},{"WeekDay":"Sat","FromTo":"10:00 – 18:00"},{"WeekDay":"Sun","FromTo":"10:00 – 18:00"}]', point(54.782313, 32.050111), 'inactive'),
       ('28c39c81-2765-439e-979d-3cfa2eea358f', 'dbb2edca-b8fb-4d9d-bb65-1461a6fed8c9', 4, 'Смоленская Крепость', 'Замки', 'ул. Барклая-Де-Толли 7, Смоленск', '+74812339585', 'museum@smolkrepost.ru', 'https://smolkrepost.ru', '[{"WeekDay":"Tue","FromTo":"10:00 – 18:00"},{"WeekDay":"Wed","FromTo":"10:00 – 18:00"},{"WeekDay":"Thu","FromTo":"10:00 – 18:00"},{"WeekDay":"Fri","FromTo":"10:00 – 17:00"},{"WeekDay":"Sat","FromTo":"10:00 – 18:00"},{"WeekDay":"Sun","FromTo":"10:00 – 18:00"}]', point(54.779037, 32.053569),'active'),
       ('24193ee8-f1ee-42d4-9052-4267ff73a760', 'afcfe1f3-2599-4a84-b814-f3bc12fb09d8', 1, 'Исторический музей', 'Музей', 'ул. Ленина, 8, Смоленск', '+74812339585', 'museum@smolkrepost.ru', 'https://smolkrepost.ru', '[{"WeekDay":"Tue","FromTo":"10:00 – 18:00"},{"WeekDay":"Wed","FromTo":"10:00 – 18:00"},{"WeekDay":"Thu","FromTo":"10:00 – 18:00"},{"WeekDay":"Fri","FromTo":"10:00 – 17:00"},{"WeekDay":"Sat","FromTo":"10:00 – 18:00"},{"WeekDay":"Sun","FromTo":"10:00 – 18:00"}]', point(54.782313, 32.050111),''),
       ('5c09dc7e-a5a2-4e94-a529-41b39ca8022d', 'afcfe1f3-2599-4a84-b814-f3bc12fb09d8', 2, 'Улица Большая Советская', 'Улица', 'Большая Советская улица, Смоленск', '+74812339585', 'museum@smolkrepost.ru', 'https://smolkrepost.ru', '[{"WeekDay":"Tue","FromTo":"10:00 – 18:00"},{"WeekDay":"Wed","FromTo":"10:00 – 18:00"},{"WeekDay":"Thu","FromTo":"10:00 – 18:00"},{"WeekDay":"Fri","FromTo":"10:00 – 17:00"},{"WeekDay":"Sat","FromTo":"10:00 – 18:00"},{"WeekDay":"Sun","FromTo":"10:00 – 18:00"}]', point(54.783627, 32.053533),'inactive');

-- +goose Down
delete from "quests_steps"
where "id" in ('51219208-8c66-4320-ac02-d37f7998ef63',
               'd297d769-cab7-4604-93a5-64101fe0df42',
               '6c35331c-9df1-4257-8b91-1be0e355bc2c',
               '28c39c81-2765-439e-979d-3cfa2eea358f',
               '24193ee8-f1ee-42d4-9052-4267ff73a760',
               '5c09dc7e-a5a2-4e94-a529-41b39ca8022d');

delete from "quests"
where "id" in ('dbb2edca-b8fb-4d9d-bb65-1461a6fed8c9', 'afcfe1f3-2599-4a84-b814-f3bc12fb09d8');

delete from "settlements"
where "id" in ('f218690b-4948-4217-8c76-7a433f533f42', 'a5e5ea94-e176-433f-9dac-b9840a039fe0');
