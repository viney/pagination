--
-- PostgreSQL database dump
--

SET statement_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;

--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET search_path = public, pg_catalog;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: user_profile; Type: TABLE; Schema: public; Owner: viney; Tablespace: 
--

DROP TABLE IF EXISTS user_profile;
CREATE TABLE user_profile (
    uid integer NOT NULL,
    name text,
    age integer
);


ALTER TABLE public.user_profile OWNER TO viney;

--
-- Data for Name: user_profile; Type: TABLE DATA; Schema: public; Owner: viney
--

COPY user_profile (uid, name, age) FROM stdin;
1	viney	32
2	viney	32
3	viney	32
4	viney	32
5	viney	32
6	viney	32
7	viney	32
8	viney	32
9	viney	32
10	viney	32
11	viney	32
12	viney	32
13	viney	32
14	viney	32
15	viney	32
16	viney	32
17	viney	32
18	viney	32
19	viney	32
20	viney	32
21	viney	32
22	viney	32
23	viney	32
24	viney	32
25	viney	32
26	viney	32
27	viney	32
28	viney	32
29	viney	32
30	viney	32
31	viney	32
32	viney	32
33	viney	32
34	viney	32
35	viney	32
36	viney	32
37	viney	32
38	viney	32
39	viney	32
40	viney	32
41	viney	32
42	viney	32
43	viney	32
44	viney	32
45	viney	32
46	viney	32
47	viney	32
48	viney	32
49	viney	32
50	viney	32
51	viney	32
52	viney	32
53	viney	32
54	viney	32
55	viney	32
56	viney	32
57	viney	32
58	viney	32
59	viney	32
60	viney	32
61	viney	32
62	viney	32
63	viney	32
64	viney	32
65	viney	32
66	viney	32
67	viney	32
68	viney	32
69	viney	32
70	viney	32
71	viney	32
72	viney	32
73	viney	32
74	viney	32
75	viney	32
76	viney	32
77	viney	32
78	viney	32
79	viney	32
80	viney	32
81	viney	32
82	viney	32
83	viney	32
84	viney	32
85	viney	32
86	viney	32
87	viney	32
88	viney	32
89	viney	32
90	viney	32
91	viney	32
92	viney	32
93	viney	32
94	viney	32
95	viney	32
96	viney	32
97	viney	32
98	viney	32
99	viney	32
100	viney	32
\.


--
-- Name: user_profile_pkey; Type: CONSTRAINT; Schema: public; Owner: viney; Tablespace: 
--

ALTER TABLE ONLY user_profile
    ADD CONSTRAINT user_profile_pkey PRIMARY KEY (uid);


--
-- Name: public; Type: ACL; Schema: -; Owner: postgres
--

REVOKE ALL ON SCHEMA public FROM PUBLIC;
REVOKE ALL ON SCHEMA public FROM postgres;
GRANT ALL ON SCHEMA public TO postgres;
GRANT ALL ON SCHEMA public TO PUBLIC;


--
-- PostgreSQL database dump complete
--

