SELECT * 
FROM exvssearchxb.airframes
JOIN exvssearchxb.airframe_costs ON exvssearchxb.airframes.airframe_cost_id = exvssearchxb.airframe_costs.id
JOIN exvssearchxb.pilots ON exvssearchxb.airframes.pilot_id = exvssearchxb.pilots.id
JOIN exvssearchxb.title_of_works ON exvssearchxb.airframes.title_of_work_id = exvssearchxb.title_of_works.id
JOIN exvssearchxb.awaken_types ON exvssearchxb.airframes.awaken_type_id = exvssearchxb.awaken_types.id
WHERE exvssearchxb.airframes.name LIKE 'ケンプファー';
