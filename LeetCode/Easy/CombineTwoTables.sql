SELECT firstName, lastName, city, state
FROM Person AS p
LEFT OUTER JOIN Address AS a ON p.personId = a.personId
