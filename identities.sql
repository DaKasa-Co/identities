--
-- PostgreSQL database dump
--

-- Dumped from database version 15.3 (Ubuntu 15.3-0ubuntu0.23.04.1)
-- Dumped by pg_dump version 15.3 (Ubuntu 15.3-0ubuntu0.23.04.1)

-- Started on 2023-07-12 12:18:18 -03

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
-- TOC entry 3 (class 3079 OID 16533)
-- Name: pgcrypto; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS pgcrypto WITH SCHEMA public;


--
-- TOC entry 3406 (class 0 OID 0)
-- Dependencies: 3
-- Name: EXTENSION pgcrypto; Type: COMMENT; Schema: -; Owner: -
--

COMMENT ON EXTENSION pgcrypto IS 'cryptographic functions';


--
-- TOC entry 2 (class 3079 OID 16511)
-- Name: uuid-ossp; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;


--
-- TOC entry 3407 (class 0 OID 0)
-- Dependencies: 2
-- Name: EXTENSION "uuid-ossp"; Type: COMMENT; Schema: -; Owner: -
--

COMMENT ON EXTENSION "uuid-ossp" IS 'generate universally unique identifiers (UUIDs)';


--
-- TOC entry 218 (class 1255 OID 16499)
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
-- TOC entry 219 (class 1255 OID 16500)
-- Name: update_at(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.update_at() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
   IF row(NEW.*) IS DISTINCT FROM row(OLD.*) THEN
      NEW.timestamp_update = now();
      RETURN NEW;
   ELSE
      RETURN OLD;
   END IF;
END;
$$;


SET default_table_access_method = heap;

--
-- TOC entry 217 (class 1259 OID 16573)
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
    timestamp_update timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    timestamp_created timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


--
-- TOC entry 216 (class 1259 OID 16572)
-- Name: users_phonenumber_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.users_phonenumber_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- TOC entry 3408 (class 0 OID 0)
-- Dependencies: 216
-- Name: users_phonenumber_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.users_phonenumber_seq OWNED BY public.users.phonenumber;


--
-- TOC entry 3252 (class 2604 OID 16577)
-- Name: users phonenumber; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.users ALTER COLUMN phonenumber SET DEFAULT nextval('public.users_phonenumber_seq'::regclass);


--
-- TOC entry 3409 (class 0 OID 0)
-- Dependencies: 216
-- Name: users_phonenumber_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.users_phonenumber_seq', 1, false);


--
-- TOC entry 3255 (class 2620 OID 16581)
-- Name: users crypt_passw; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER crypt_passw BEFORE INSERT ON public.users FOR EACH ROW EXECUTE FUNCTION public.encrypt_passwords();


--
-- TOC entry 3256 (class 2620 OID 16582)
-- Name: users update_timestamp; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER update_timestamp BEFORE UPDATE ON public.users FOR EACH ROW EXECUTE FUNCTION public.update_at();


-- Completed on 2023-07-12 12:18:18 -03

--
-- PostgreSQL database dump complete
--

