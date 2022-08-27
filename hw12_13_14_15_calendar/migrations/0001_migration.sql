-- +goose Up
CREATE TABLE public.calendar_event (
    id serial NOT NULL,
    user_id int NOT NULL,
    title varchar NOT NULL,
    datetime_event timestamp NOT NULL,
    duration_event interval NOT NULL,
    description varchar NULL,
    notification_interval interval NULL,
    CONSTRAINT calendar_event_pk PRIMARY KEY (id)
);
CREATE INDEX calendar_event_user_id_idx ON public.calendar_event (user_id);
CREATE INDEX calendar_event_datetime_event_idx ON public.calendar_event (datetime_event);

-- +goose Down
DROP TABLE public.calendar_event;