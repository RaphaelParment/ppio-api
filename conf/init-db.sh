#!/bin/bash

echo "Creation of the database structure."
cockroach sql --insecure < ../sql/structure.sql

echo "Inserting dummy data."
ppio -initDbData

echo "Ready to develop..."
