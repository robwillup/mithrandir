# AI-900: Fundamental Principles of Machine Learning on Azure

This course is all about Machine Learning on Azure and how to use it to make predictions, classify things and group items together.

Topics:

- What is machine Learning
- Predict a numeric value with regression
- Predicting categories or classes
- Utilizing Azure tools for machine learning

## Machine Learning

### An introduction to Machine Learning

> Technique that uses mathematics and statistics to create a model that can predict unknown values

Predicting a numerical value with Machine Learning is an example of regression.

To predict the price of a car we can use historical data.

* This historical data is called a dataset

**Model:**

> A Machine Learning model is a file that has been trained with data and an algorithm to recognize patterns.
> Once the model's been trained it can be deployed to ingest data and start making predictions.

In the case example, when using previous data, consider the following example:

|Engine size|Gas Mileage|Price|
|:----------|:----------|:----|
|3400|18|3000|
|2000|40|5000|

If we're using `Engine Size` and `Gas Mileage` to predict `Price` then `Engine Size` and `Gas Mileage`
are called `Features` and `Price` is called `Label`.

Data -> Machine Learning Platform -> Train a model -> Deploy the Model

#### Azure Services to support Machine Learning

* Azure Automated Machine Learning
  * Quick start with machine learning. Select a dataset, pick a label, off it goes.
* Azure Machine Learning Designer
  * Create a pipeline to train a machine learning model, without writing code, using a drag-and-drop interface


