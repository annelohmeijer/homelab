import xgboost as xgb
from sklearn.preprocessing import LabelEncoder
import numpy as np
import pandas as pd


def load_model(path: str) -> xgb.Booster:
    bst = xgb.Booster()
    bst.load_model(path)
    return bst


def load_data(path: str) -> pd.DataFrame:
    return pd.read_csv(path, parse_dates=['date_recorded'], index_col='id')


def preprocess(df: pd.DataFrame, bst: xgb.Booster, sample_size=100) -> xgb.DMatrix:
    le_object = LabelEncoder()
    label_encode_me_please = ['funder', 'installer', 'wpt_name', 'basin', 'subvillage', 'region', 'lga', 'ward',
                              'public_meeting', 'recorded_by', 'scheme_management', 'scheme_name', 'permit',
                              'extraction_type', 'extraction_type_group', 'extraction_type_class', 'management',
                              'management_group', 'payment', 'payment_type', 'water_quality', 'quality_group',
                              'quantity', 'quantity_group', 'source', 'source_type', 'source_class', 'waterpoint_type',
                              'waterpoint_type_group']
    data_transformed = df[label_encode_me_please].apply(le_object.fit_transform)
    df[label_encode_me_please] = data_transformed
    df['date_recorded'] = df.date_recorded.astype(int) / 10 ** 9
    cols_when_model_builds = bst.feature_names
    df = df[cols_when_model_builds]
    # Take a random sample of the data, normally this would be some unseen new record that has to be predicted
    # For the assignment sampling the training data is sufficient
    sample_data = df.sample(sample_size)

    return xgb.DMatrix(sample_data)


def predict(xgb_matrix: xgb.DMatrix, bst: xgb.Booster) -> list:
    y_prob = bst.predict(xgb_matrix)
    y_pred = np.argmax(y_prob, axis=1)
    class_names = ['functional', 'functional needs repair', 'non functional']
    return [class_names[x] for x in y_pred]




if __name__ == '__main__':

    model = load_model('files/model.json')

    data = load_data('files/water_pump_set.csv')

    matrix = preprocess(data, model)

    predictions = predict(matrix, model)

    print('predictions: ', predictions)
