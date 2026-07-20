TRUNCATE orders, customers, jobs, documents RESTART IDENTITY CASCADE;
INSERT INTO customers(email, region) SELECT 'user'||g||'@example.test', CASE WHEN g < 90000 THEN 'RU' ELSE 'OTHER' END FROM generate_series(1,100000) g;
INSERT INTO orders(customer_id,status,total_cents,created_at) SELECT 1 + (g % 100000), (ARRAY['new','paid','done'])[1+(g%3)], 100+(g%50000), now()-(g%365)*interval '1 day' FROM generate_series(1,300000) g;
INSERT INTO jobs(payload) SELECT jsonb_build_object('n',g) FROM generate_series(1,1000) g;
ANALYZE;
