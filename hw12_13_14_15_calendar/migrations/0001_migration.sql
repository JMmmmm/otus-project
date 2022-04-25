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

insert into public.calendar_event (user_id, title, datetime_event, duration_event, description, notification_interval)
VALUES (1, 'test', '2004-10-19 10:23:54', '3 days 04:05:06', 'test test', '3 days 04:05:06');

-- +goose Down
DROP TABLE public.calendar_event;