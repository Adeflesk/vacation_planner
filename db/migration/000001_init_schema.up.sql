CREATE TABLE "tripplan" (
  "id" bigserial PRIMARY KEY,
  "trip_name" varchar NOT NULL,
  "start_date" date NOT NULL,
  "end_date" date NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "countries" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "continent_name" varchar NOT NULL
);

CREATE TABLE "locations" (
  "id" bigserial PRIMARY KEY,
  "location_name" varchar NOT NULL,
  "location_description" varchar NOT NULL,
  "country_id" bigint NOT NULL
);

CREATE TABLE "activities" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "activity_type" bigint NOT NULL,
  "description" varchar NOT NULL,
  "time_allocated" time,
  "location_id" bigint NOT NULL
);

CREATE TABLE "activity_type" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL
);

CREATE TABLE "car_rental" (
  "id" bigserial PRIMARY KEY,
  "Company" varchar NOT NULL,
  "cost" bigint NOT NULL,
  "booking_ref" varchar NOT NULL,
  "website" varchar NOT NULL,
  "pickupdate" timestamp NOT NULL,
  "dropoffdate" timestamp NOT NULL,
  "pickuplocation" bigint NOT NULL,
  "dropoflocation" bigint NOT NULL,
  "trip_id" bigint NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "flights" (
  "id" bigserial PRIMARY KEY,
  "bookingid" varchar NOT NULL,
  "airline_name" varchar NOT NULL,
  "airport" varchar NOT NULL,
  "cost" bigint NOT NULL,
  "originLocationId" int NOT NULL,
  "destinationId" int NOT NULL,
  "trip_id" bigint NOT NULL,
  "departure_date" timestamp NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "accommodation" (
  "id" bigserial PRIMARY KEY,
  "accommodation_name" varchar NOT NULL,
  "pernight" bigint NOT NULL,
  "accommodation_type" bigint NOT NULL,
  "accommodation_description" varchar NOT NULL,
  "webaddress" varchar NOT NULL,
  "emailaddress" varchar NOT NULL,
  "phonenumber" varchar NOT NULL,
  "area" bigint NOT NULL
);

CREATE TABLE "accommodation_type" (
  "id" bigserial PRIMARY KEY,
  "type" varchar NOT NULL
);

CREATE TABLE "food" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "area" bigint NOT NULL,
  "food_type" bigint NOT NULL,
  "webaddress" varchar NOT NULL
);

CREATE TABLE "food_type" (
  "id" bigserial PRIMARY KEY,
  "type" varchar NOT NULL
);

CREATE TABLE "food_plan" (
  "id" bigserial PRIMARY KEY,
  "foodId" bigint NOT NULL,
  "tripId" bigint NOT NULL,
  "date" timestamp NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "accommmodation_nights" (
  "id" bigserial PRIMARY KEY,
  "accommodation" bigint NOT NULL,
  "tripId" bigInt NOT NULL,
  "date" date NOT NULL,
  "reservationNumber" varchar NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "activity_days" (
  "id" bigserial PRIMARY KEY,
  "activityId" bigint NOT NULL,
  "tripplanId" bigint NOT NULL,
  "date" date NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

ALTER TABLE "locations" ADD FOREIGN KEY ("country_id") REFERENCES "countries" ("id");

ALTER TABLE "activities" ADD FOREIGN KEY ("activity_type") REFERENCES "activity_type" ("id");

ALTER TABLE "activities" ADD FOREIGN KEY ("location_id") REFERENCES "locations" ("id");

ALTER TABLE "car_rental" ADD FOREIGN KEY ("pickuplocation") REFERENCES "locations" ("id");

ALTER TABLE "car_rental" ADD FOREIGN KEY ("dropoflocation") REFERENCES "locations" ("id");

ALTER TABLE "car_rental" ADD FOREIGN KEY ("trip_id") REFERENCES "tripplan" ("id");

ALTER TABLE "flights" ADD FOREIGN KEY ("originLocationId") REFERENCES "locations" ("id");

ALTER TABLE "flights" ADD FOREIGN KEY ("destinationId") REFERENCES "locations" ("id");

ALTER TABLE "flights" ADD FOREIGN KEY ("trip_id") REFERENCES "tripplan" ("id");

ALTER TABLE "accommodation" ADD FOREIGN KEY ("accommodation_type") REFERENCES "accommodation_type" ("id");

ALTER TABLE "accommodation" ADD FOREIGN KEY ("area") REFERENCES "locations" ("id");

ALTER TABLE "food" ADD FOREIGN KEY ("area") REFERENCES "locations" ("id");

ALTER TABLE "food" ADD FOREIGN KEY ("food_type") REFERENCES "food_type" ("id");

ALTER TABLE "food_plan" ADD FOREIGN KEY ("foodId") REFERENCES "food" ("id");

ALTER TABLE "food_plan" ADD FOREIGN KEY ("tripId") REFERENCES "tripplan" ("id");

ALTER TABLE "accommmodation_nights" ADD FOREIGN KEY ("accommodation") REFERENCES "accommodation" ("id");

ALTER TABLE "accommmodation_nights" ADD FOREIGN KEY ("tripId") REFERENCES "tripplan" ("id");

ALTER TABLE "activity_days" ADD FOREIGN KEY ("activityId") REFERENCES "activities" ("id");

ALTER TABLE "activity_days" ADD FOREIGN KEY ("tripplanId") REFERENCES "tripplan" ("id");
