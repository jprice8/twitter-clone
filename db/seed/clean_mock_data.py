#%%
import pandas as pd

#%%
df = pd.read_csv('~/data/amazon/mock_amazon_data.csv')

#%%
class MockData:
    """
    Class to represent the Amazon mock data.
    """
    def load(self) -> pd.DataFrame:
        """
        Load in the csv from local storage.
        """
        pass