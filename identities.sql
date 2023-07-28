--
-- PostgreSQL database dump
--

-- Dumped from database version 15.3 (Ubuntu 15.3-0ubuntu0.23.04.1)
-- Dumped by pg_dump version 15.3 (Ubuntu 15.3-0ubuntu0.23.04.1)

-- Started on 2023-07-28 18:27:51 -03

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

--
-- TOC entry 2 (class 3079 OID 24796)
-- Name: pgcrypto; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS pgcrypto WITH SCHEMA public;


--
-- TOC entry 3419 (class 0 OID 0)
-- Dependencies: 2
-- Name: EXTENSION pgcrypto; Type: COMMENT; Schema: -; Owner: -
--

COMMENT ON EXTENSION pgcrypto IS 'cryptographic functions';


--
-- TOC entry 3 (class 3079 OID 24833)
-- Name: uuid-ossp; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;


--
-- TOC entry 3420 (class 0 OID 0)
-- Dependencies: 3
-- Name: EXTENSION "uuid-ossp"; Type: COMMENT; Schema: -; Owner: -
--

COMMENT ON EXTENSION "uuid-ossp" IS 'generate universally unique identifiers (UUIDs)';


--
-- TOC entry 265 (class 1255 OID 24844)
-- Name: add_date_interval(interval); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.add_date_interval(add_interval interval) RETURNS timestamp without time zone
    LANGUAGE plpgsql STRICT
    AS $$
BEGIN
   RETURN CURRENT_TIMESTAMP + add_interval;
END;
$$;


--
-- TOC entry 266 (class 1255 OID 24845)
-- Name: add_ten_minutes(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.add_ten_minutes() RETURNS trigger
    LANGUAGE plpgsql
    AS $$BEGIN
	NEW.expire_at = add_date_interval('10 minutes');
	RETURN NEW;
END;$$;


--
-- TOC entry 267 (class 1255 OID 24846)
-- Name: encrypt_passwords(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.encrypt_passwords() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
	NEW.password = crypt(LTRIM(NEW.password), gen_salt('bf'));
    RETURN NEW;
END;
$$;


--
-- TOC entry 268 (class 1255 OID 24847)
-- Name: random_between(integer, integer); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.random_between(low integer, high integer) RETURNS integer
    LANGUAGE plpgsql STRICT
    AS $$
BEGIN
   RETURN floor(random()* (high-low + 1) + low);
END;
$$;


--
-- TOC entry 269 (class 1255 OID 24848)
-- Name: updated_at(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.updated_at() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
   IF row(NEW.*) IS DISTINCT FROM row(OLD.*) THEN
      NEW.updated_at = now();
      RETURN NEW;
   ELSE
      RETURN OLD;
   END IF;
END;
$$;


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 216 (class 1259 OID 24849)
-- Name: recovery; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.recovery (
    id uuid NOT NULL,
    user_id uuid NOT NULL,
    validation integer NOT NULL,
    expire_at timestamp with time zone NOT NULL
);


--
-- TOC entry 217 (class 1259 OID 24852)
-- Name: users; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.users (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    name character varying(80) NOT NULL,
    username character varying(35) NOT NULL,
    email character varying(225) NOT NULL,
    password character varying(225) NOT NULL,
    birthday date NOT NULL,
    phonenumber bigint NOT NULL,
    address character varying(225),
    picture character varying(150),
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


--
-- TOC entry 218 (class 1259 OID 24860)
-- Name: users_phonenumber_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.users_phonenumber_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- TOC entry 3421 (class 0 OID 0)
-- Dependencies: 218
-- Name: users_phonenumber_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.users_phonenumber_seq OWNED BY public.users.phonenumber;


--
-- TOC entry 3259 (class 2604 OID 24861)
-- Name: users phonenumber; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.users ALTER COLUMN phonenumber SET DEFAULT nextval('public.users_phonenumber_seq'::regclass);


--
-- TOC entry 3411 (class 0 OID 24849)
-- Dependencies: 216
-- Data for Name: recovery; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.recovery (id, user_id, validation, expire_at) FROM stdin;
\.


--
-- TOC entry 3412 (class 0 OID 24852)
-- Dependencies: 217
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.users (id, name, username, email, password, birthday, phonenumber, address, picture, updated_at, created_at) FROM stdin;
\.


--
-- TOC entry 3422 (class 0 OID 0)
-- Dependencies: 218
-- Name: users_phonenumber_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.users_phonenumber_seq', 1, false);


--
-- TOC entry 3263 (class 2606 OID 24868)
-- Name: recovery recovery_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.recovery
    ADD CONSTRAINT recovery_pkey PRIMARY KEY (id);


--
-- TOC entry 3265 (class 2606 OID 24870)
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id, email, phonenumber, username);


--
-- TOC entry 3266 (class 2620 OID 24862)
-- Name: recovery add_expire_date; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER add_expire_date BEFORE INSERT ON public.recovery FOR EACH ROW EXECUTE FUNCTION public.add_ten_minutes();


--
-- TOC entry 3267 (class 2620 OID 24865)
-- Name: users crypt_sensitive_data; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER crypt_sensitive_data BEFORE INSERT OR UPDATE OF password ON public.users FOR EACH ROW EXECUTE FUNCTION public.encrypt_passwords();


--
-- TOC entry 3268 (class 2620 OID 24866)
-- Name: users update_date; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER update_date BEFORE INSERT OR UPDATE OF id, username, name, email, password, birthday, phonenumber, address, picture, updated_at, created_at ON public.users FOR EACH ROW EXECUTE FUNCTION public.updated_at();


-- Completed on 2023-07-28 18:27:51 -03

--
-- PostgreSQL database dump complete
--

