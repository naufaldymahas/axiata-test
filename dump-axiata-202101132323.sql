PGDMP     )                     y            axiata    12.2    12.2     ~           0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false                       0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false            �           0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false            �           1262    29326    axiata    DATABASE     d   CREATE DATABASE axiata WITH TEMPLATE = template0 ENCODING = 'UTF8' LC_COLLATE = 'C' LC_CTYPE = 'C';
    DROP DATABASE axiata;
                naufaldymahas    false                        2615    2200    public    SCHEMA        CREATE SCHEMA public;
    DROP SCHEMA public;
                naufaldymahas    false            �           0    0    SCHEMA public    COMMENT     6   COMMENT ON SCHEMA public IS 'standard public schema';
                   naufaldymahas    false    3            �            1259    29363    employee_id_seq    SEQUENCE     }   CREATE SEQUENCE public.employee_id_seq
    START WITH 200001
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 &   DROP SEQUENCE public.employee_id_seq;
       public          dev    false    3            �            1259    29361    employee_year_seq    SEQUENCE     }   CREATE SEQUENCE public.employee_year_seq
    START WITH 2020
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 (   DROP SEQUENCE public.employee_year_seq;
       public          dev    false    3            �            1259    29459    users    TABLE     4  CREATE TABLE public.users (
    user_id bigint NOT NULL,
    user_name text NOT NULL,
    password text NOT NULL,
    employee_id text NOT NULL,
    birth_place text NOT NULL,
    birth_date date NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL
);
    DROP TABLE public.users;
       public         heap    dev    false    3            �            1259    29457    users_user_id_seq    SEQUENCE     z   CREATE SEQUENCE public.users_user_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 (   DROP SEQUENCE public.users_user_id_seq;
       public          dev    false    3    205            �           0    0    users_user_id_seq    SEQUENCE OWNED BY     G   ALTER SEQUENCE public.users_user_id_seq OWNED BY public.users.user_id;
          public          dev    false    204            �           2604    29462    users user_id    DEFAULT     n   ALTER TABLE ONLY public.users ALTER COLUMN user_id SET DEFAULT nextval('public.users_user_id_seq'::regclass);
 <   ALTER TABLE public.users ALTER COLUMN user_id DROP DEFAULT;
       public          dev    false    204    205    205            �           2606    29471    users users_employee_id_key 
   CONSTRAINT     ]   ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_employee_id_key UNIQUE (employee_id);
 E   ALTER TABLE ONLY public.users DROP CONSTRAINT users_employee_id_key;
       public            dev    false    205            �           2606    29467    users users_pkey 
   CONSTRAINT     S   ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (user_id);
 :   ALTER TABLE ONLY public.users DROP CONSTRAINT users_pkey;
       public            dev    false    205            �           2606    29469    users users_user_name_key 
   CONSTRAINT     Y   ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_user_name_key UNIQUE (user_name);
 C   ALTER TABLE ONLY public.users DROP CONSTRAINT users_user_name_key;
       public            dev    false    205           