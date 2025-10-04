--
-- PostgreSQL database dump
--

\restrict OeR8nqdMABHwYfSeLF8dOM4r7kk7W2EwsRKpPETJUdpmos5Rr04h1t4tVPo1i4d

-- Dumped from database version 14.19 (Homebrew)
-- Dumped by pg_dump version 14.19 (Homebrew)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: schema_migration; Type: TABLE; Schema: public; Owner: navneetshukla
--

CREATE TABLE public.schema_migration (
    version character varying(14) NOT NULL
);


ALTER TABLE public.schema_migration OWNER TO navneetshukla;

--
-- Name: users; Type: TABLE; Schema: public; Owner: navneetshukla
--

CREATE TABLE public.users (
    name character varying(100),
    email character varying(100),
    password character varying(100),
    mobile character varying(100)
);


ALTER TABLE public.users OWNER TO navneetshukla;

--
-- Name: schema_migration schema_migration_pkey; Type: CONSTRAINT; Schema: public; Owner: navneetshukla
--

ALTER TABLE ONLY public.schema_migration
    ADD CONSTRAINT schema_migration_pkey PRIMARY KEY (version);


--
-- Name: schema_migration_version_idx; Type: INDEX; Schema: public; Owner: navneetshukla
--

CREATE UNIQUE INDEX schema_migration_version_idx ON public.schema_migration USING btree (version);


--
-- PostgreSQL database dump complete
--

\unrestrict OeR8nqdMABHwYfSeLF8dOM4r7kk7W2EwsRKpPETJUdpmos5Rr04h1t4tVPo1i4d

