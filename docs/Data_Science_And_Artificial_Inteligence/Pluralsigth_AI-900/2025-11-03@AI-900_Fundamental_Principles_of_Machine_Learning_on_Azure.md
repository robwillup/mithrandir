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

### Types of Machine Learning

#### Regression

* Use historical data to predict a numerical value
* Finding the price of a car based on its Features
* The historical data will include features and labels
* Predicting a numerical value is called regression

#### Classification

A form of machine learning that is used to predict which category or class an item belongs to

##### Classification Types

* Binary Classification - looking for one of two options
* Multi-class classification - multiple possible options

Both Regression and Classification use historical data with known features and labels to trains the machine learning
model. This is known as **Supervised Learning**.

#### Clustering

Group similar items into clusters based on their features. An example is grouping flowers based on petal count,
petal size, leaf size, etc.

This is **unsupervised learning**

#### Supervised vs unsupervised

|Supervised|Unsupervised|
|:---------|:-----------|
|The data includes known features and label| The data has features but no known labels|
|Trying to predict something| Trying to group similar items|

#### Deep Learning

Uses a structure of artificial neural networks. This consists on multiple types of inputs, outputs, and even hidden layers.

Deep learning allows a machine to train itself!

#### Machine Learning vs Deep Learning

| Machine Learning | Deep Learning |
|:-----------------|:--------------|
| Can use small amounts of data | Need to use a large amount of data |
| Works on low-end machines | Requires high-end machines |
| Requires features to be identified by the end user| Learns features and can create new features |
| Trains quickly - minutes, hours | Slow to train |
| Outputs are (usually) one format - numeric or class| Output can have multiple formats |

##### Deep Learning Use cases

* Identifying patterns in unstructured data
* Image caption generation
* Automatic text and speech translation


