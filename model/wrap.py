"""Wrap entrypoint for mlflow that gets trained model and stores it as mlflow artifacts in mlflow run."""
import logging
import pickle
import sys

import mlflow
import numpy as np
import pandas as pd
import xgboost as xgb
from sklearn.preprocessing import LabelEncoder

logging.basicConfig(stream=sys.stdout, level=logging.DEBUG)


class ModelWrapper(mlflow.pyfunc.PythonModel):
    """Model wrapper that handles incoming data on model side when served."""

    def __init__(self):
        self.model = None
        self._logger = logging.getLogger(__name__)
        self.return_columns = ["id", "pred", "pred_proba"]

    def load_context(self, context: mlflow.pyfunc.PythonModelContext) -> None:
        """Loading models from mlflow model artifact"""
        self.model = pickle.load(open(context.artifacts["model"], "rb"))
        self._logger.info(f"return_columns: {self.return_columns}")

    def predict(
        self, context: mlflow.pyfunc.PythonModelContext, df: pd.DataFrame
    ) -> dict:
        """Predict endpoint of model."""
        self._logger.info(f"Input dataframe size {df.shape}")
        self._logger.info(f"Input dataframe.head(): {df.head()}")
        matrix = self.preprocess(df)
        y_prob = self.model.predict(matrix)

        # get predicted class
        y_pred = np.argmax(y_prob, axis=1)
        class_names = ["functional", "functional needs repair", "non functional"]
        predictions = [class_names[x] for x in y_pred]

        # get probabilites, useful for lineage and/or thresholding
        y_pred_probas = np.max(y_prob, axis=1)

        df = df.assign(pred=predictions, pred_proba=y_pred_probas)

        return df[self.return_columns].to_dict(orient="records")

    def preprocess(self, df: pd.DataFrame) -> xgb.DMatrix:
        """Preprocess scoring data."""
        le_object = LabelEncoder()
        label_encode_me_please = [
            "funder",
            "installer",
            "wpt_name",
            "basin",
            "subvillage",
            "region",
            "lga",
            "ward",
            "public_meeting",
            "recorded_by",
            "scheme_management",
            "scheme_name",
            "permit",
            "extraction_type",
            "extraction_type_group",
            "extraction_type_class",
            "management",
            "management_group",
            "payment",
            "payment_type",
            "water_quality",
            "quality_group",
            "quantity",
            "quantity_group",
            "source",
            "source_type",
            "source_class",
            "waterpoint_type",
            "waterpoint_type_group",
        ]
        data_transformed = df[label_encode_me_please].apply(le_object.fit_transform)
        df[label_encode_me_please] = data_transformed
        df["date_recorded"] = df.date_recorded.astype(int) / 10**9
        cols_when_model_builds = self.model.feature_names
        df = df[cols_when_model_builds]
        return xgb.DMatrix(df)


if __name__ == "__main__":
    with mlflow.start_run():

        model_obj = xgb.Booster()
        model_obj.load_model("files/model.json")

        pickle.dump(model_obj, open("model.pkl", "wb"))

        artifacts = {"model": "model.pkl"}

        model_wrapper = ModelWrapper()

        mlflow.pyfunc.log_model(
            artifact_path="",
            python_model=model_wrapper,
            conda_env="mlflow_conda_env.yml",
            artifacts=artifacts,
        )
