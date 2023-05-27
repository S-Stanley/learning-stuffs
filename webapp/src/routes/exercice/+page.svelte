<script lang="ts">
    import { answer } from './store';

    const data = [
       {
            question: "etre",
            answer: "ser"
       },
       {
            question: "avetre (etre situé)",
            answer: "estar"
       },

       {
            question: "avoir",
            answer: "tener"
       },
       {
            question: "faire",
            answer: "hacer"
       },
       {
            question: "aller",
            answer: "ir"
       },
       {
            question: "pouvoir",
            answer: "poder"
       },
       {
            question: "savoir",
            answer: "saber"
       },
       {
            question: "mettre",
            answer: "poner"
       },
       {
            question: "dire",
            answer: "decir"
       },
       {
            question: "vouloir",
            answer: "querer"
       },
       {
            question: "parler",
            answer: "hablar"
       },
       {
            question: "donner",
            answer: "dar"
       },
       {
            question: "voir (regarder)",
            answer: "ver"
       },
       {
            question: "manger",
            answer: "comer"
       },
       {
            question: "prendre",
            answer: "tomar"
       },
       {
            question: "vivre",
            answer: "vivir"
       },
       {
            question: "avoir besoin de",
            answer: "necessitar"
       },
       {
            question: "rester",
            answer: "quedar"
       },
       {
            question: "venir",
            answer: "venir"
       },
    ];

    function random(): number {
        return Math.floor(Math.random() * data.length);
    }

    let question_number = random();
    let is_finish = false;
    let points = 0;
    let number_of_question_answered = 0;
    const number_of_questions_to_ask = 10;

    function onSubmit(){
        if ($answer.input === data[question_number].answer) {
            points++;
        } else {
            alert(`bad, the answer was ${data[question_number].answer}`);
        }
        if (number_of_question_answered == number_of_questions_to_ask -1) {
            is_finish = true;
        } else {
            question_number = random();
            (document.querySelector('#input-user-answer') as HTMLElement)?.focus();
        }
        $answer.input = "";
        number_of_question_answered++;
    }

    function try_again(){
        points = 0;
        question_number = 0;
        is_finish = false;
        number_of_question_answered = 0;
    }
</script>

{#if is_finish === true}
    <p>You finished the course with {points} success answer on {number_of_questions_to_ask}</p>
    <button on:click={try_again}>Try again</button>
{:else}
    <h1>{number_of_question_answered + 1}/{number_of_questions_to_ask} Traduisez en espagnol: {data[question_number].question}</h1>
    <input type="text" bind:value={$answer.input} id="input-user-answer"/>
    <button on:click={onSubmit}>OK</button>
{/if}