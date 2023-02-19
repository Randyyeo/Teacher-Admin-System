# Bank Account

Simple banking system that handles operations on bank accounts. At the moment, the project only is capable of three features:
- Depositing an amount
- Withdrawing an amount
- Printing the account statement

All user inputs will be done through the command line.

## Design Architecture

The application is designed using an MVC design pattern. The view portion comes from the `App.py` which is where the application is running from. By the user's interaction with the View, different handlers (Controller layer) will be called. For example, if the user enters 'd' into the command line, the `DepositHandler` will be called where the business logic lies within it. It then updates the Model components, a.k.a. User, by depositing or withdrawing money.

In the `src` folder, it contains 4 folders:
1. Constants -> containing all the constant values
2. Handlers -> contain the different handlers to handle different business logic
3. Models -> contain enum classes as well as our User class to maintain the states of the user
4. Utils -> Different util functions to handle certain common functions

## Launch virtual environment

To ensure this application can run on all operating systems, the instructions below will showcase how it can be ran on a virtual environment. To create a virtual environment for this project, run the following command:
```
# MacOS
python3 -m venv venv
# Windows/Linux
python -m venv venv
```

This will create a `venv` folder where you can launch the virtual environment. To launch it:
```
source venv/bin/activate
```

You should see a (venv) in your command line. To exit the virtual environment, run: 
```
deactivate
```

## Installation

To install the dependencies for the project, run the following command:

```
pip install -r requirements.txt
```

## Usage

To run the project, use the following command in the terminal:
```
python src/app.py
```

## Tests

To run the unit tests provided:
```
# There are 3 folders to run the tests

pytest tests/handlers_tests
pytest tests/models_tests
pytest tests/utils_tests
```

## Assumptions

There are a few assumptions in this project:
1. It will give an error if an invalid input such as "A" is entered into the commmand line when the user is prompted for an action.
2. When entering an amount for the deposit/withdraw, the user cannot enter invalid types such as letters as well as numbers with more than 2 decimal points. It will throw an error and ask the user to select an action again.

