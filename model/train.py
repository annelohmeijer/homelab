import pandas as pd
from sklearn.model_selection import train_test_split
import xgboost as xgb
from sklearn import preprocessing
import numpy as np
import pickle

if __name__ == '__main__':
    labels = pd.read_csv('../files/water_pump_labels.csv', index_col='id')
    df = pd.read_csv('../files/water_pump_set.csv', parse_dates=['date_recorded'], index_col='id')

    label_encode_me_please = ['funder', 'installer', 'wpt_name', 'basin', 'subvillage', 'region', 'lga', 'ward',
                              'public_meeting', 'recorded_by', 'scheme_management', 'scheme_name', 'permit',
                              'extraction_type', 'extraction_type_group', 'extraction_type_class', 'management',
                              'management_group', 'payment', 'payment_type', 'water_quality', 'quality_group',
                              'quantity', 'quantity_group', 'source', 'source_type', 'source_class', 'waterpoint_type',
                              'waterpoint_type_group']
    le_object = preprocessing.LabelEncoder()
    df_joined = df.join(labels)
    transformed = df_joined[label_encode_me_please].apply(le_object.fit_transform)
    df_joined[label_encode_me_please] = transformed
    df_joined['date_recorded'] = df_joined.date_recorded.astype(int) / 10 ** 9

    le_target = preprocessing.LabelEncoder()

    target_encoded = df_joined[['status_group']].apply(le_target.fit_transform)
    columns = list(set(df_joined.columns) - {'status_group'})
    x_train, x_test, y_train, y_test = train_test_split(df_joined[columns], target_encoded)

    params = {
        'objective': 'multi:softprob',
        'num_class': 3,
        'min_child_weight': 15,
        'eta': 0.3,
        'gamma': 5,
        'max_depth': 5,
        'nthread': 3,
        'eval_metric': 'mlogloss',
    }
    dtrain = xgb.DMatrix(x_train, label=y_train)
    dtest = xgb.DMatrix(x_test)
    model = xgb.train(params, dtrain, num_boost_round=50, verbose_eval=5)

    model.save_model('../files/model.json')

    with open('../files/le_target.pkl', 'wb') as f:
        pickle.dump(le_target, f)

    y_prob = model.predict(dtest)
    from sklearn.metrics import log_loss, confusion_matrix

    print('Log-Loss on test is {}'.format(log_loss(y_test, y_prob)))
    y_pred = np.argmax(y_prob * np.array([.3, 1, 1]), axis=1)

    class_names = list(le_target.inverse_transform([0, 1, 2]))
    matrix = confusion_matrix(y_test, y_pred)
    print(matrix)