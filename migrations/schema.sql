--
-- PostgreSQL database dump
--

-- Dumped from database version 16.1
-- Dumped by pg_dump version 16.1

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
-- Name: admin; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.admin (
    id integer NOT NULL,
    first_name character varying(255) DEFAULT ''::character varying NOT NULL,
    last_name character varying(255) DEFAULT ''::character varying NOT NULL,
    thumbnail character varying(255) DEFAULT ''::character varying NOT NULL,
    email character varying(255) NOT NULL,
    password character varying(60) NOT NULL,
    access_level integer DEFAULT 1 NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.admin OWNER TO postgres;

--
-- Name: admin_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.admin_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.admin_id_seq OWNER TO postgres;

--
-- Name: admin_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.admin_id_seq OWNED BY public.admin.id;


--
-- Name: schema_migration; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.schema_migration (
    version character varying(14) NOT NULL
);


ALTER TABLE public.schema_migration OWNER TO postgres;

--
-- Name: service; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.service (
    id integer NOT NULL,
    name character varying(255) DEFAULT ''::character varying NOT NULL,
    available integer DEFAULT 1 NOT NULL,
    minimum_charge integer DEFAULT 0 NOT NULL,
    category_id integer NOT NULL,
    sub_category_id integer NOT NULL,
    description character varying(255) DEFAULT ''::character varying NOT NULL,
    thumbnail character varying(255) DEFAULT ''::character varying NOT NULL,
    country character varying(255) DEFAULT 'Bangladesh'::character varying NOT NULL,
    division character varying(255) DEFAULT ''::character varying NOT NULL,
    district character varying(255) DEFAULT ''::character varying NOT NULL,
    city character varying(255) DEFAULT ''::character varying NOT NULL,
    created_at date NOT NULL,
    updated_at date NOT NULL
);


ALTER TABLE public.service OWNER TO postgres;

--
-- Name: service_category_main; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.service_category_main (
    id integer NOT NULL,
    name character varying(255) DEFAULT ''::character varying NOT NULL,
    available integer DEFAULT 1 NOT NULL,
    description character varying(255) DEFAULT ''::character varying NOT NULL,
    thumbnail character varying(255) DEFAULT ''::character varying NOT NULL,
    created_at date NOT NULL,
    updated_at date NOT NULL
);


ALTER TABLE public.service_category_main OWNER TO postgres;

--
-- Name: service_category_main_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.service_category_main_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.service_category_main_id_seq OWNER TO postgres;

--
-- Name: service_category_main_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.service_category_main_id_seq OWNED BY public.service_category_main.id;


--
-- Name: service_category_sub; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.service_category_sub (
    id integer NOT NULL,
    name character varying(255) DEFAULT ''::character varying NOT NULL,
    available integer DEFAULT 1 NOT NULL,
    category_id integer NOT NULL,
    description character varying(255) DEFAULT ''::character varying NOT NULL,
    thumbnail character varying(255) DEFAULT ''::character varying NOT NULL,
    created_at date NOT NULL,
    updated_at date NOT NULL
);


ALTER TABLE public.service_category_sub OWNER TO postgres;

--
-- Name: service_category_sub_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.service_category_sub_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.service_category_sub_id_seq OWNER TO postgres;

--
-- Name: service_category_sub_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.service_category_sub_id_seq OWNED BY public.service_category_sub.id;


--
-- Name: service_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.service_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.service_id_seq OWNER TO postgres;

--
-- Name: service_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.service_id_seq OWNED BY public.service.id;


--
-- Name: admin id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.admin ALTER COLUMN id SET DEFAULT nextval('public.admin_id_seq'::regclass);


--
-- Name: service id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.service ALTER COLUMN id SET DEFAULT nextval('public.service_id_seq'::regclass);


--
-- Name: service_category_main id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.service_category_main ALTER COLUMN id SET DEFAULT nextval('public.service_category_main_id_seq'::regclass);


--
-- Name: service_category_sub id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.service_category_sub ALTER COLUMN id SET DEFAULT nextval('public.service_category_sub_id_seq'::regclass);


--
-- Name: admin admin_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.admin
    ADD CONSTRAINT admin_pkey PRIMARY KEY (id);


--
-- Name: schema_migration schema_migration_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.schema_migration
    ADD CONSTRAINT schema_migration_pkey PRIMARY KEY (version);


--
-- Name: service_category_main service_category_main_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.service_category_main
    ADD CONSTRAINT service_category_main_pkey PRIMARY KEY (id);


--
-- Name: service_category_sub service_category_sub_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.service_category_sub
    ADD CONSTRAINT service_category_sub_pkey PRIMARY KEY (id);


--
-- Name: service service_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.service
    ADD CONSTRAINT service_pkey PRIMARY KEY (id);


--
-- Name: schema_migration_version_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX schema_migration_version_idx ON public.schema_migration USING btree (version);


--
-- Name: service service_category_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.service
    ADD CONSTRAINT service_category_id_fkey FOREIGN KEY (category_id) REFERENCES public.service_category_main(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: service_category_sub service_category_sub_category_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.service_category_sub
    ADD CONSTRAINT service_category_sub_category_id_fkey FOREIGN KEY (category_id) REFERENCES public.service_category_main(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

