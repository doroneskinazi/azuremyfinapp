select tm.id, tm.date, tm.amount, tc.name, tt.type
from 	myfp.trans_money tm
inner join myfp.trans_category tc on tm.category_id = tc.id
inner join myfp.trans_type tt on tm.type_id = tt.id



