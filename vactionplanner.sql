CREATE TABLE "tripplan" (
  "id" bigserial PRIMARY KEY,
  "trip_name" varchar,
  "start_date" date,
  "end_date" date,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "countries" (
  "id" int PRIMARY KEY,
  "name" varchar,
  "continent_name" varchar
);

CREATE TABLE "locations" (
  "id" bigserial PRIMARY KEY,
  "location_name" varchar,
  "location_description" varchar,
  "country_id" int NOT NULL
);

CREATE TABLE "activities" (
  "id" bigserial PRIMARY KEY,
  "name" varchar,
  "activitytype" bigint NOT NULL,
  "description" varchar,
  "time_allocated" time,
  "location_id" bigint NOT NULL
);

CREATE TABLE "activityType" (
  "id" bigserial PRIMARY KEY,
  "name" varchar
);

CREATE TABLE "car_rental" (
  "id" bigserial PRIMARY KEY,
  "Company" varchar,
  "cost" bigint,
  "booking_ref" varchar,
  "website" varchar,
  "pickupdate" timestamp,
  "dropoffdate" timestamp,
  "pickuplocation" bigint NOT NULL,
  "dropoflocation" bigint NOT NULL,
  "trip_id" bigint NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "flights" (
  "id" bigserial PRIMARY KEY,
  "bookingid" varchar,
  "airline_name" varchar,
  "airport" varchar,
  "cost" bigint,
  "originLocationId" int NOT NULL,
  "destinationId" int NOT NULL,
  "trip_id" bigint NOT NULL,
  "departure_date" timestamp,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "accommodation" (
  "id" bigserial PRIMARY KEY,
  "accommodation_name" varchar,
  "pernight" bigint,
  "type" bigint NOT NULL,
  "description" varchar,
  "emailaddress" varchar,
  "phonenumber" varchar,
  "locationId" bigint NOT NULL
);

CREATE TABLE "accommodationType" (
  "id" bigserial PRIMARY KEY,
  "type" varchar
);

CREATE TABLE "food" (
  "id" bigserial PRIMARY KEY,
  "name" varchar,
  "locationId" bigint NOT NULL,
  "foodtype" bigint NOT NULL,
  "webaddress" varchar
);

CREATE TABLE "foodType" (
  "id" bigserial PRIMARY KEY,
  "type" varchar
);

CREATE TABLE "foodPlan" (
  "id" bigserial PRIMARY KEY,
  "foodId" bigint NOT NULL,
  "tripId" bigint NOT NULL,
  "date" timestamp,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "AccommmodationNights" (
  "id" bigserial PRIMARY KEY,
  "accommodation" bigint NOT NULL,
  "tripId" bigInt NOT NULL,
  "date" date,
  "reservationNumber" varchar,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "activityDays" (
  "id" bigserial PRIMARY KEY,
  "activityId" bigint NOT NULL,
  "tripplanId" bigint NOT NULL,
  "date" date,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

ALTER TABLE "locations" ADD FOREIGN KEY ("country_id") REFERENCES "countries" ("id");

ALTER TABLE "activities" ADD FOREIGN KEY ("activitytype") REFERENCES "activityType" ("id");

ALTER TABLE "activities" ADD FOREIGN KEY ("location_id") REFERENCES "locations" ("id");

ALTER TABLE "car_rental" ADD FOREIGN KEY ("pickuplocation") REFERENCES "locations" ("id");

ALTER TABLE "car_rental" ADD FOREIGN KEY ("dropoflocation") REFERENCES "locations" ("id");

ALTER TABLE "car_rental" ADD FOREIGN KEY ("trip_id") REFERENCES "tripplan" ("id");

ALTER TABLE "flights" ADD FOREIGN KEY ("originLocationId") REFERENCES "locations" ("id");

ALTER TABLE "flights" ADD FOREIGN KEY ("destinationId") REFERENCES "locations" ("id");

ALTER TABLE "flights" ADD FOREIGN KEY ("trip_id") REFERENCES "tripplan" ("id");

ALTER TABLE "accommodation" ADD FOREIGN KEY ("type") REFERENCES "accommodationType" ("id");

ALTER TABLE "accommodation" ADD FOREIGN KEY ("locationId") REFERENCES "locations" ("id");

ALTER TABLE "food" ADD FOREIGN KEY ("locationId") REFERENCES "locations" ("id");

ALTER TABLE "food" ADD FOREIGN KEY ("foodtype") REFERENCES "foodType" ("id");

ALTER TABLE "foodPlan" ADD FOREIGN KEY ("foodId") REFERENCES "food" ("id");

ALTER TABLE "foodPlan" ADD FOREIGN KEY ("tripId") REFERENCES "tripplan" ("id");

ALTER TABLE "AccommmodationNights" ADD FOREIGN KEY ("accommodation") REFERENCES "accommodation" ("id");

ALTER TABLE "AccommmodationNights" ADD FOREIGN KEY ("tripId") REFERENCES "tripplan" ("id");

ALTER TABLE "activityDays" ADD FOREIGN KEY ("activityId") REFERENCES "activities" ("id");

ALTER TABLE "activityDays" ADD FOREIGN KEY ("tripplanId") REFERENCES "tripplan" ("id");
