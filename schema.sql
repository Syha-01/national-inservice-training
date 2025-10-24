--
-- PostgreSQL database dump
--

\restrict 3W8Cl7YIGSVRG3KTqV0ATBW5gltvNcVkgQqraAT9xBQgkqO6Ts8S06cwo7K21cA

-- Dumped from database version 17.6 (Ubuntu 17.6-1.pgdg24.04+1)
-- Dumped by pg_dump version 17.6 (Ubuntu 17.6-1.pgdg24.04+1)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: citext; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS citext WITH SCHEMA public;


--
-- Name: EXTENSION citext; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION citext IS 'data type for case-insensitive character strings';


--
-- Name: update_updated_at_column(); Type: FUNCTION; Schema: public; Owner: nits
--

CREATE FUNCTION public.update_updated_at_column() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$;


ALTER FUNCTION public.update_updated_at_column() OWNER TO nits;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: course_ratings; Type: TABLE; Schema: public; Owner: nits
--

CREATE TABLE public.course_ratings (
    id integer NOT NULL,
    session_enrollment_id integer NOT NULL,
    score integer NOT NULL,
    comment text,
    created_at timestamp with time zone DEFAULT now(),
    CONSTRAINT course_ratings_score_check CHECK (((score >= 1) AND (score <= 5)))
);


ALTER TABLE public.course_ratings OWNER TO nits;

--
-- Name: course_ratings_id_seq; Type: SEQUENCE; Schema: public; Owner: nits
--

CREATE SEQUENCE public.course_ratings_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.course_ratings_id_seq OWNER TO nits;

--
-- Name: course_ratings_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: nits
--

ALTER SEQUENCE public.course_ratings_id_seq OWNED BY public.course_ratings.id;


--
-- Name: courses; Type: TABLE; Schema: public; Owner: nits
--

CREATE TABLE public.courses (
    id integer NOT NULL,
    title character varying(255) NOT NULL,
    description text,
    category character varying(20) NOT NULL,
    credit_hours numeric(5,2) NOT NULL,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    CONSTRAINT courses_category_check CHECK (((category)::text = ANY ((ARRAY['Mandatory'::character varying, 'Elective'::character varying])::text[]))),
    CONSTRAINT courses_credit_hours_check CHECK ((credit_hours > (0)::numeric))
);


ALTER TABLE public.courses OWNER TO nits;

--
-- Name: courses_id_seq; Type: SEQUENCE; Schema: public; Owner: nits
--

CREATE SEQUENCE public.courses_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.courses_id_seq OWNER TO nits;

--
-- Name: courses_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: nits
--

ALTER SEQUENCE public.courses_id_seq OWNED BY public.courses.id;


--
-- Name: facilitator_ratings; Type: TABLE; Schema: public; Owner: nits
--

CREATE TABLE public.facilitator_ratings (
    id integer NOT NULL,
    session_enrollment_id integer NOT NULL,
    facilitator_id integer NOT NULL,
    score integer NOT NULL,
    comment text,
    created_at timestamp with time zone DEFAULT now(),
    CONSTRAINT facilitator_ratings_score_check CHECK (((score >= 1) AND (score <= 5)))
);


ALTER TABLE public.facilitator_ratings OWNER TO nits;

--
-- Name: facilitator_ratings_id_seq; Type: SEQUENCE; Schema: public; Owner: nits
--

CREATE SEQUENCE public.facilitator_ratings_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.facilitator_ratings_id_seq OWNER TO nits;

--
-- Name: facilitator_ratings_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: nits
--

ALTER SEQUENCE public.facilitator_ratings_id_seq OWNED BY public.facilitator_ratings.id;


--
-- Name: facilitators; Type: TABLE; Schema: public; Owner: nits
--

CREATE TABLE public.facilitators (
    id integer NOT NULL,
    first_name character varying(100) NOT NULL,
    last_name character varying(100) NOT NULL,
    email character varying(255),
    personnel_id integer,
    version integer DEFAULT 1 NOT NULL
);


ALTER TABLE public.facilitators OWNER TO nits;

--
-- Name: facilitators_id_seq; Type: SEQUENCE; Schema: public; Owner: nits
--

CREATE SEQUENCE public.facilitators_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.facilitators_id_seq OWNER TO nits;

--
-- Name: facilitators_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: nits
--

ALTER SEQUENCE public.facilitators_id_seq OWNED BY public.facilitators.id;


--
-- Name: formations; Type: TABLE; Schema: public; Owner: nits
--

CREATE TABLE public.formations (
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    region_id integer NOT NULL
);


ALTER TABLE public.formations OWNER TO nits;

--
-- Name: formations_id_seq; Type: SEQUENCE; Schema: public; Owner: nits
--

CREATE SEQUENCE public.formations_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.formations_id_seq OWNER TO nits;

--
-- Name: formations_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: nits
--

ALTER SEQUENCE public.formations_id_seq OWNED BY public.formations.id;


--
-- Name: permissions; Type: TABLE; Schema: public; Owner: nits
--

CREATE TABLE public.permissions (
    id integer NOT NULL,
    code text NOT NULL,
    description text
);


ALTER TABLE public.permissions OWNER TO nits;

--
-- Name: permissions_id_seq; Type: SEQUENCE; Schema: public; Owner: nits
--

CREATE SEQUENCE public.permissions_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.permissions_id_seq OWNER TO nits;

--
-- Name: permissions_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: nits
--

ALTER SEQUENCE public.permissions_id_seq OWNED BY public.permissions.id;


--
-- Name: personnel; Type: TABLE; Schema: public; Owner: nits
--

CREATE TABLE public.personnel (
    id integer NOT NULL,
    regulation_number character varying(50) NOT NULL,
    first_name character varying(100) NOT NULL,
    last_name character varying(100) NOT NULL,
    sex character varying(10) NOT NULL,
    rank_id integer,
    formation_id integer,
    posting_id integer,
    is_active boolean DEFAULT true,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    CONSTRAINT personnel_sex_check CHECK (((sex)::text = ANY ((ARRAY['Male'::character varying, 'Female'::character varying])::text[])))
);


ALTER TABLE public.personnel OWNER TO nits;

--
-- Name: personnel_id_seq; Type: SEQUENCE; Schema: public; Owner: nits
--

CREATE SEQUENCE public.personnel_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.personnel_id_seq OWNER TO nits;

--
-- Name: personnel_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: nits
--

ALTER SEQUENCE public.personnel_id_seq OWNED BY public.personnel.id;


--
-- Name: postings; Type: TABLE; Schema: public; Owner: nits
--

CREATE TABLE public.postings (
    id integer NOT NULL,
    name character varying(255) NOT NULL
);


ALTER TABLE public.postings OWNER TO nits;

--
-- Name: postings_id_seq; Type: SEQUENCE; Schema: public; Owner: nits
--

CREATE SEQUENCE public.postings_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.postings_id_seq OWNER TO nits;

--
-- Name: postings_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: nits
--

ALTER SEQUENCE public.postings_id_seq OWNED BY public.postings.id;


--
-- Name: ranks; Type: TABLE; Schema: public; Owner: nits
--

CREATE TABLE public.ranks (
    id integer NOT NULL,
    name character varying(100) NOT NULL,
    abbreviation character varying(20)
);


ALTER TABLE public.ranks OWNER TO nits;

--
-- Name: ranks_id_seq; Type: SEQUENCE; Schema: public; Owner: nits
--

CREATE SEQUENCE public.ranks_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.ranks_id_seq OWNER TO nits;

--
-- Name: ranks_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: nits
--

ALTER SEQUENCE public.ranks_id_seq OWNED BY public.ranks.id;


--
-- Name: regions; Type: TABLE; Schema: public; Owner: nits
--

CREATE TABLE public.regions (
    id integer NOT NULL,
    name character varying(100) NOT NULL
);


ALTER TABLE public.regions OWNER TO nits;

--
-- Name: regions_id_seq; Type: SEQUENCE; Schema: public; Owner: nits
--

CREATE SEQUENCE public.regions_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.regions_id_seq OWNER TO nits;

--
-- Name: regions_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: nits
--

ALTER SEQUENCE public.regions_id_seq OWNED BY public.regions.id;


--
-- Name: role_permissions; Type: TABLE; Schema: public; Owner: nits
--

CREATE TABLE public.role_permissions (
    role_id integer NOT NULL,
    permission_id integer NOT NULL
);


ALTER TABLE public.role_permissions OWNER TO nits;

--
-- Name: roles; Type: TABLE; Schema: public; Owner: nits
--

CREATE TABLE public.roles (
    id integer NOT NULL,
    name character varying(50) NOT NULL
);


ALTER TABLE public.roles OWNER TO nits;

--
-- Name: roles_id_seq; Type: SEQUENCE; Schema: public; Owner: nits
--

CREATE SEQUENCE public.roles_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.roles_id_seq OWNER TO nits;

--
-- Name: roles_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: nits
--

ALTER SEQUENCE public.roles_id_seq OWNED BY public.roles.id;


--
-- Name: schema_migrations; Type: TABLE; Schema: public; Owner: nits
--

CREATE TABLE public.schema_migrations (
    version bigint NOT NULL,
    dirty boolean NOT NULL
);


ALTER TABLE public.schema_migrations OWNER TO nits;

--
-- Name: session_enrollment; Type: TABLE; Schema: public; Owner: nits
--

CREATE TABLE public.session_enrollment (
    id integer NOT NULL,
    personnel_id integer NOT NULL,
    session_id integer NOT NULL,
    completion_date date,
    status character varying(50) DEFAULT 'Enrolled'::character varying,
    created_at timestamp with time zone DEFAULT now(),
    CONSTRAINT session_enrollment_status_check CHECK (((status)::text = ANY ((ARRAY['Enrolled'::character varying, 'Completed'::character varying, 'Failed'::character varying, 'Withdrew'::character varying])::text[])))
);


ALTER TABLE public.session_enrollment OWNER TO nits;

--
-- Name: session_enrollment_id_seq; Type: SEQUENCE; Schema: public; Owner: nits
--

CREATE SEQUENCE public.session_enrollment_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.session_enrollment_id_seq OWNER TO nits;

--
-- Name: session_enrollment_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: nits
--

ALTER SEQUENCE public.session_enrollment_id_seq OWNED BY public.session_enrollment.id;


--
-- Name: session_facilitators; Type: TABLE; Schema: public; Owner: nits
--

CREATE TABLE public.session_facilitators (
    session_id integer NOT NULL,
    facilitator_id integer NOT NULL
);


ALTER TABLE public.session_facilitators OWNER TO nits;

--
-- Name: tokens; Type: TABLE; Schema: public; Owner: nits
--

CREATE TABLE public.tokens (
    hash bytea NOT NULL,
    user_id bigint NOT NULL,
    expiry timestamp with time zone NOT NULL,
    scope text NOT NULL
);


ALTER TABLE public.tokens OWNER TO nits;

--
-- Name: training_sessions; Type: TABLE; Schema: public; Owner: nits
--

CREATE TABLE public.training_sessions (
    id integer NOT NULL,
    course_id integer NOT NULL,
    start_date date NOT NULL,
    end_date date NOT NULL,
    location character varying(255),
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    version integer DEFAULT 1 NOT NULL
);


ALTER TABLE public.training_sessions OWNER TO nits;

--
-- Name: training_sessions_id_seq; Type: SEQUENCE; Schema: public; Owner: nits
--

CREATE SEQUENCE public.training_sessions_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.training_sessions_id_seq OWNER TO nits;

--
-- Name: training_sessions_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: nits
--

ALTER SEQUENCE public.training_sessions_id_seq OWNED BY public.training_sessions.id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: nits
--

CREATE TABLE public.users (
    id integer NOT NULL,
    email character varying(255) NOT NULL,
    password_hash character varying(255) NOT NULL,
    role_id integer NOT NULL,
    personnel_id integer,
    activated boolean DEFAULT true,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now()
);


ALTER TABLE public.users OWNER TO nits;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: nits
--

CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.users_id_seq OWNER TO nits;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: nits
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: users_permissions; Type: TABLE; Schema: public; Owner: nits
--

CREATE TABLE public.users_permissions (
    user_id bigint NOT NULL,
    permission_id bigint NOT NULL
);


ALTER TABLE public.users_permissions OWNER TO nits;

--
-- Name: course_ratings id; Type: DEFAULT; Schema: public; Owner: nits
--

ALTER TABLE ONLY public.course_ratings ALTER COLUMN id SET DEFAULT nextval('public.course_ratings_id_seq'::regclass);


--
-- Name: courses id; Type: DEFAULT; Schema: public; Owner: nits
--

ALTER TABLE ONLY public.courses ALTER COLUMN id SET DEFAULT nextval('public.courses_id_seq'::regclass);


--
-- Name: facilitator_ratings id; Type: DEFAULT; Schema: public; Owner: nits
--

ALTER TABLE ONLY public.facilitator_ratings ALTER COLUMN id SET DEFAULT nextval('public.facilitator_ratings_id_seq'::regclass);


--
-- Name: facilitators id; Type: DEFAULT; Schema: public; Owner: nits
--

ALTER TABLE ONLY public.facilitators ALTER COLUMN id SET DEFAULT nextval('public.facilitators_id_seq'::regclass);


--
-- Name: formations id; Type: DEFAULT; Schema: public; Owner: nits
--

ALTER TABLE ONLY public.formations ALTER COLUMN id SET DEFAULT nextval('public.formations_id_seq'::regclass);


--
-- Name: permissions id; Type: DEFAULT; Schema: public; Owner: nits
--

ALTER TABLE ONLY public.permissions ALTER COLUMN id SET DEFAULT nextval('public.permissions_id_seq'::regclass);


--
-- Name: personnel id; Type: DEFAULT; Schema: public; Owner: nits
--

ALTER TABLE ONLY public.personnel ALTER COLUMN id SET DEFAULT nextval('public.personnel_id_seq'::regclass);


--
-- Name: postings id; Type: DEFAULT; Schema: public; Owner: nits
--

ALTER TABLE ONLY public.postings ALTER COLUMN id SET DEFAULT nextval('public.postings_id_seq'::regclass);


--
-- Name: ranks id; Type: DEFAULT; Schema: public; Owner: nits
--

ALTER TABLE ONLY public.ranks ALTER COLUMN id SET DEFAULT nextval('public.ranks_id_seq'::regclass);


--
-- Name: regions id; Type: DEFAULT; Schema: public; Owner: nits
--

ALTER TABLE ONLY public.regions ALTER COLUMN id SET DEFAULT nextval('public.regions_id_seq'::regclass);


--
-- Name: roles id; Type: DEFAULT; Schema: public; Owner: nits
--

ALTER TABLE ONLY public.roles ALTER COLUMN id SET DEFAULT nextval('public.roles_id_seq'::regclass);


--
-- Name: session_enrollment id; Type: DEFAULT; Schema: public; Owner: nits
--

ALTER TABLE ONLY public.session_enrollment ALTER COLUMN id SET DEFAULT nextval('public.session_enrollment_id_seq'::regclass);


--
-- Name: training_sessions id; Type: DEFAULT; Schema: public; Owner: nits
--

ALTER TABLE ONLY public.training_sessions ALTER COLUMN id SET DEFAULT nextval('public.training_sessions_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: nits
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Name: course_ratings course_ratings_pkey; Type: CONSTRAINT; Schema: public; Owner: nits
--

ALTER TABLE ONLY public.course_ratings
    ADD CONSTRAINT course_ratings_pkey PRIMARY KEY (id);


--
-- Name: course_ratings course_ratings_session_enrollment_id_key; Type: CONSTRAINT; Schema: public; Owner: nits
--

ALTER TABLE ONLY public.course_ratings
    ADD CONSTRAINT course_ratings_session_enrollment_id_key UNIQUE (session_enrollment_id);


--
-- Name: courses courses_pkey; Type: CONSTRAINT; Schema: public; Owner: nits
--

ALTER TABLE ONLY public.courses
    ADD CONSTRAINT courses_pkey PRIMARY KEY (id);


--
-- Name: facilitator_ratings facilitator_ratings_pkey; Type: CONSTRAINT; Schema: public; Owner: nits
--

ALTER TABLE ONLY public.facilitator_ratings
    ADD CONSTRAINT facilitator_ratings_pkey PRIMARY KEY (id);


--
-- Name: facilitator_ratings facilitator_ratings_session_enrollment_id_facilitator_id_key; Type: CONSTRAINT; Schema: public; Owner: nits
--

ALTER TABLE ONLY public.facilitator_ratings
    ADD CONSTRAINT facilitator_ratings_session_enrollment_id_facilitator_id_key UNIQUE (session_enrollment_id, facilitator_id);


--
-- Name: facilitators facilitators_email_key; Type: CONSTRAINT; Schema: public; Owner: nits
--

ALTER TABLE ONLY public.facilitators
    ADD CONSTRAINT facilitators_email_key UNIQUE (email);


--
-- Name: facilitators facilitators_personnel_id_key; Type: CONSTRAINT; Schema: public; Owner: nits
--

ALTER TABLE ONLY public.facilitators
    ADD CONSTRAINT facilitators_personnel_id_key UNIQUE (personnel_id);


--
-- Name: facilitators facilitators_pkey; Type: CONSTRAINT; Schema: public; Owner: nits
--

ALTER TABLE ONLY public.facilitators
    ADD CONSTRAINT facilitators_pkey PRIMARY KEY (id);


--
-- Name: formations formations_name_key; Type: CONSTRAINT; Schema: public; Owner: nits
--

ALTER TABLE ONLY public.formations
    ADD CONSTRAINT formations_name_key UNIQUE (name);


--
-- Name: formations formations_pkey; Type: CONSTRAINT; Schema: public; Owner: nits
--

ALTER TABLE ONLY public.formations
    ADD CONSTRAINT formations_pkey PRIMARY KEY (id);


--
-- Name: permissions permissions_code_key; Type: CONSTRAINT; Schema: public; Owner: nits
--

ALTER TABLE ONLY public.permissions
    ADD CONSTRAINT permissions_code_key UNIQUE (code);


--
-- Name: permissions permissions_pkey; Type: CONSTRAINT; Schema: public; Owner: nits
--

ALTER TABLE ONLY public.permissions
    ADD CONSTRAINT permissions_pkey PRIMARY KEY (id);


--
-- Name: personnel personnel_pkey; Type: CONSTRAINT; Schema: public; Owner: nits
--

ALTER TABLE ONLY public.personnel
    ADD CONSTRAINT personnel_pkey PRIMARY KEY (id);


--
-- Name: personnel personnel_regulation_number_key; Type: CONSTRAINT; Schema: public; Owner: nits
--

ALTER TABLE ONLY public.personnel
    ADD CONSTRAINT personnel_regulation_number_key UNIQUE (regulation_number);


--
-- Name: postings postings_name_key; Type: CONSTRAINT; Schema: public; Owner: nits
--

ALTER TABLE ONLY public.postings
    ADD CONSTRAINT postings_name_key UNIQUE (name);


--
-- Name: postings postings_pkey; Type: CONSTRAINT; Schema: public; Owner: nits
--

ALTER TABLE ONLY public.postings
    ADD CONSTRAINT postings_pkey PRIMARY KEY (id);


--
-- Name: ranks ranks_abbreviation_key; Type: CONSTRAINT; Schema: public; Owner: nits
--

ALTER TABLE ONLY public.ranks
    ADD CONSTRAINT ranks_abbreviation_key UNIQUE (abbreviation);


--
-- Name: ranks ranks_name_key; Type: CONSTRAINT; Schema: public; Owner: nits
--

ALTER TABLE ONLY public.ranks
    ADD CONSTRAINT ranks_name_key UNIQUE (name);


--
-- Name: ranks ranks_pkey; Type: CONSTRAINT; Schema: public; Owner: nits
--

ALTER TABLE ONLY public.ranks
    ADD CONSTRAINT ranks_pkey PRIMARY KEY (id);


--
-- Name: regions regions_name_key; Type: CONSTRAINT; Schema: public; Owner: nits
--

ALTER TABLE ONLY public.regions
    ADD CONSTRAINT regions_name_key UNIQUE (name);


--
-- Name: regions regions_pkey; Type: CONSTRAINT; Schema: public; Owner: nits
--

ALTER TABLE ONLY public.regions
    ADD CONSTRAINT regions_pkey PRIMARY KEY (id);


--
-- Name: role_permissions role_permissions_pkey; Type: CONSTRAINT; Schema: public; Owner: nits
--

ALTER TABLE ONLY public.role_permissions
    ADD CONSTRAINT role_permissions_pkey PRIMARY KEY (role_id, permission_id);


--
-- Name: roles roles_name_key; Type: CONSTRAINT; Schema: public; Owner: nits
--

ALTER TABLE ONLY public.roles
    ADD CONSTRAINT roles_name_key UNIQUE (name);


--
-- Name: roles roles_pkey; Type: CONSTRAINT; Schema: public; Owner: nits
--

ALTER TABLE ONLY public.roles
    ADD CONSTRAINT roles_pkey PRIMARY KEY (id);


--
-- Name: schema_migrations schema_migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: nits
--

ALTER TABLE ONLY public.schema_migrations
    ADD CONSTRAINT schema_migrations_pkey PRIMARY KEY (version);


--
-- Name: session_enrollment session_enrollment_personnel_id_session_id_key; Type: CONSTRAINT; Schema: public; Owner: nits
--

ALTER TABLE ONLY public.session_enrollment
    ADD CONSTRAINT session_enrollment_personnel_id_session_id_key UNIQUE (personnel_id, session_id);


--
-- Name: session_enrollment session_enrollment_pkey; Type: CONSTRAINT; Schema: public; Owner: nits
--

ALTER TABLE ONLY public.session_enrollment
    ADD CONSTRAINT session_enrollment_pkey PRIMARY KEY (id);


--
-- Name: session_facilitators session_facilitators_pkey; Type: CONSTRAINT; Schema: public; Owner: nits
--

ALTER TABLE ONLY public.session_facilitators
    ADD CONSTRAINT session_facilitators_pkey PRIMARY KEY (session_id, facilitator_id);


--
-- Name: tokens tokens_pkey; Type: CONSTRAINT; Schema: public; Owner: nits
--

ALTER TABLE ONLY public.tokens
    ADD CONSTRAINT tokens_pkey PRIMARY KEY (hash);


--
-- Name: training_sessions training_sessions_pkey; Type: CONSTRAINT; Schema: public; Owner: nits
--

ALTER TABLE ONLY public.training_sessions
    ADD CONSTRAINT training_sessions_pkey PRIMARY KEY (id);


--
-- Name: users users_email_key; Type: CONSTRAINT; Schema: public; Owner: nits
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_email_key UNIQUE (email);


--
-- Name: users_permissions users_permissions_pkey; Type: CONSTRAINT; Schema: public; Owner: nits
--

ALTER TABLE ONLY public.users_permissions
    ADD CONSTRAINT users_permissions_pkey PRIMARY KEY (user_id, permission_id);


--
-- Name: users users_personnel_id_key; Type: CONSTRAINT; Schema: public; Owner: nits
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_personnel_id_key UNIQUE (personnel_id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: nits
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: idx_tokens_expiry; Type: INDEX; Schema: public; Owner: nits
--

CREATE INDEX idx_tokens_expiry ON public.tokens USING btree (expiry);


--
-- Name: idx_tokens_user_id; Type: INDEX; Schema: public; Owner: nits
--

CREATE INDEX idx_tokens_user_id ON public.tokens USING btree (user_id);


--
-- Name: users update_users_updated_at; Type: TRIGGER; Schema: public; Owner: nits
--

CREATE TRIGGER update_users_updated_at BEFORE UPDATE ON public.users FOR EACH ROW EXECUTE FUNCTION public.update_updated_at_column();


--
-- Name: course_ratings course_ratings_session_enrollment_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: nits
--

ALTER TABLE ONLY public.course_ratings
    ADD CONSTRAINT course_ratings_session_enrollment_id_fkey FOREIGN KEY (session_enrollment_id) REFERENCES public.session_enrollment(id) ON DELETE CASCADE;


--
-- Name: facilitator_ratings facilitator_ratings_facilitator_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: nits
--

ALTER TABLE ONLY public.facilitator_ratings
    ADD CONSTRAINT facilitator_ratings_facilitator_id_fkey FOREIGN KEY (facilitator_id) REFERENCES public.facilitators(id) ON DELETE CASCADE;


--
-- Name: facilitator_ratings facilitator_ratings_session_enrollment_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: nits
--

ALTER TABLE ONLY public.facilitator_ratings
    ADD CONSTRAINT facilitator_ratings_session_enrollment_id_fkey FOREIGN KEY (session_enrollment_id) REFERENCES public.session_enrollment(id) ON DELETE CASCADE;


--
-- Name: facilitators facilitators_personnel_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: nits
--

ALTER TABLE ONLY public.facilitators
    ADD CONSTRAINT facilitators_personnel_id_fkey FOREIGN KEY (personnel_id) REFERENCES public.personnel(id) ON DELETE SET NULL;


--
-- Name: formations formations_region_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: nits
--

ALTER TABLE ONLY public.formations
    ADD CONSTRAINT formations_region_id_fkey FOREIGN KEY (region_id) REFERENCES public.regions(id) ON DELETE RESTRICT;


--
-- Name: personnel personnel_formation_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: nits
--

ALTER TABLE ONLY public.personnel
    ADD CONSTRAINT personnel_formation_id_fkey FOREIGN KEY (formation_id) REFERENCES public.formations(id) ON DELETE SET NULL;


--
-- Name: personnel personnel_posting_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: nits
--

ALTER TABLE ONLY public.personnel
    ADD CONSTRAINT personnel_posting_id_fkey FOREIGN KEY (posting_id) REFERENCES public.postings(id) ON DELETE SET NULL;


--
-- Name: personnel personnel_rank_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: nits
--

ALTER TABLE ONLY public.personnel
    ADD CONSTRAINT personnel_rank_id_fkey FOREIGN KEY (rank_id) REFERENCES public.ranks(id) ON DELETE SET NULL;


--
-- Name: role_permissions role_permissions_permission_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: nits
--

ALTER TABLE ONLY public.role_permissions
    ADD CONSTRAINT role_permissions_permission_id_fkey FOREIGN KEY (permission_id) REFERENCES public.permissions(id) ON DELETE CASCADE;


--
-- Name: role_permissions role_permissions_role_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: nits
--

ALTER TABLE ONLY public.role_permissions
    ADD CONSTRAINT role_permissions_role_id_fkey FOREIGN KEY (role_id) REFERENCES public.roles(id) ON DELETE CASCADE;


--
-- Name: session_enrollment session_enrollment_personnel_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: nits
--

ALTER TABLE ONLY public.session_enrollment
    ADD CONSTRAINT session_enrollment_personnel_id_fkey FOREIGN KEY (personnel_id) REFERENCES public.personnel(id) ON DELETE CASCADE;


--
-- Name: session_enrollment session_enrollment_session_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: nits
--

ALTER TABLE ONLY public.session_enrollment
    ADD CONSTRAINT session_enrollment_session_id_fkey FOREIGN KEY (session_id) REFERENCES public.training_sessions(id) ON DELETE CASCADE;


--
-- Name: session_facilitators session_facilitators_facilitator_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: nits
--

ALTER TABLE ONLY public.session_facilitators
    ADD CONSTRAINT session_facilitators_facilitator_id_fkey FOREIGN KEY (facilitator_id) REFERENCES public.facilitators(id) ON DELETE CASCADE;


--
-- Name: session_facilitators session_facilitators_session_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: nits
--

ALTER TABLE ONLY public.session_facilitators
    ADD CONSTRAINT session_facilitators_session_id_fkey FOREIGN KEY (session_id) REFERENCES public.training_sessions(id) ON DELETE CASCADE;


--
-- Name: tokens tokens_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: nits
--

ALTER TABLE ONLY public.tokens
    ADD CONSTRAINT tokens_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- Name: training_sessions training_sessions_course_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: nits
--

ALTER TABLE ONLY public.training_sessions
    ADD CONSTRAINT training_sessions_course_id_fkey FOREIGN KEY (course_id) REFERENCES public.courses(id) ON DELETE CASCADE;


--
-- Name: users_permissions users_permissions_permission_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: nits
--

ALTER TABLE ONLY public.users_permissions
    ADD CONSTRAINT users_permissions_permission_id_fkey FOREIGN KEY (permission_id) REFERENCES public.permissions(id) ON DELETE CASCADE;


--
-- Name: users_permissions users_permissions_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: nits
--

ALTER TABLE ONLY public.users_permissions
    ADD CONSTRAINT users_permissions_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- Name: users users_personnel_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: nits
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_personnel_id_fkey FOREIGN KEY (personnel_id) REFERENCES public.personnel(id) ON DELETE SET NULL;


--
-- Name: users users_role_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: nits
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_role_id_fkey FOREIGN KEY (role_id) REFERENCES public.roles(id) ON DELETE RESTRICT;


--
-- PostgreSQL database dump complete
--

\unrestrict 3W8Cl7YIGSVRG3KTqV0ATBW5gltvNcVkgQqraAT9xBQgkqO6Ts8S06cwo7K21cA

